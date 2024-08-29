package main

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/tokenchar"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var esClient *elasticsearch.TypedClient

const (
	productIndex = "products"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 5 * time.Second,
			DialContext:           (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
		},
		Username: "elastic",
		Password: "elastic",
	}

	var err error
	esClient, err = elasticsearch.NewTypedClient(cfg)
	if err != nil {
		slog.Error("Error creating es client:", err)
		os.Exit(1)
	}

	//if err = deleteProductIndex(); err != nil {
	//	slog.Error("Error deleting product index:", err)
	//}
	//
	//if err = createProductIndex(); err != nil {
	//	slog.Error("Error creating product index:", err)
	//	os.Exit(1)
	//}
	//
	//if err = bulkAddProduct(); err != nil {
	//	slog.Error("Error bulkAddProduct:", err)
	//	os.Exit(1)
	//}

	// get product with price between 1000 to 6000
	products, err := filterProductsByPrice(1000, 100000)
	if err != nil {
		slog.Error("Error filterProductsByPrice:", err)
	}
	slog.Info("List Products", slog.Any("Products", products))

	// search ssu
	const searchTerm = "SSu"
	products2, err := searchProductByName(searchTerm)
	if err != nil {
		slog.Error("Error searchProductByName:", err)
	}
	slog.Info("List Products with searchProductByName", slog.Any("search term", searchTerm), slog.Any("Products", products2))

	// search ssu with price between 1000 to 10000
	products3, err := searchAndFilterProduct(searchTerm, 1000, 10000)
	if err != nil {
		slog.Error("Error searchAndFilterProduct:", err)
	}
	slog.Info("List Products with searchAndFilterProduct ", slog.Any("search term", searchTerm), slog.Any("Products", products3))

	// search ssu with highlight
	products4, err := searchProductWithHighlight(searchTerm)
	if err != nil {
		slog.Error("Error searchProductWithHighlight:", err)
	}
	slog.Info("List Products with searchProductWithHighlight", slog.Any("search term", searchTerm), slog.Any("Products", products4))

}

func deleteProductIndex() error {
	_, err := esClient.Indices.Delete(productIndex).IgnoreUnavailable(true).Do(nil)
	if err != nil {
		return err
	}
	return nil

}

func createProductIndex() error {
	analyzer := "ngram_analyzer"
	maxNGramDiff := 10
	settings := &types.IndexSettings{
		MaxNgramDiff: &maxNGramDiff,
		Analysis: &types.IndexSettingsAnalysis{
			Analyzer: map[string]types.Analyzer{
				"ngram_analyzer": types.CustomAnalyzer{
					Tokenizer: "ngram_tokenizer",
					Filter:    []string{"lowercase"},
				},
			},
			Tokenizer: map[string]types.Tokenizer{
				"ngram_tokenizer": types.NGramTokenizer{
					Type:       "ngram",
					MinGram:    1,
					MaxGram:    5,
					TokenChars: []tokenchar.TokenChar{tokenchar.Letter, tokenchar.Digit},
				},
			},
		},
	}

	indexTrue := true
	mappings := &types.TypeMapping{
		Properties: map[string]types.Property{
			"name": types.TextProperty{
				Type:     "text",
				Index:    &indexTrue,
				Analyzer: &analyzer,
			},
			"price": types.IntegerNumberProperty{
				Type: "integer",
			},
		},
	}

	request := &create.Request{
		Settings: settings,
		Mappings: mappings,
	}
	_, err := esClient.Indices.
		Create(productIndex).
		Request(request).
		Do(nil)
	if err != nil {
		return err
	}
	return nil
}

func bulkAddProduct() error {
	products := []Product{
		{Name: "Susu ABC", Price: 10000},
		{Name: "Cabai Merah", Price: 200},
		{Name: "Susu Dancow", Price: 45000},
		{Name: "Sayur Kol", Price: 5800},
	}
	for _, product := range products {
		_, err := esClient.Index(productIndex).Request(product).Do(nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func filterProductsByPrice(from, to int) ([]Product, error) {
	gte := types.Float64(from)
	lte := types.Float64(to)
	res, err := esClient.Search().Index(productIndex).
		Request(&search.Request{
			Query: &types.Query{
				Range: map[string]types.RangeQuery{
					"price": types.NumberRangeQuery{
						Gte: &gte,
						Lte: &lte,
					},
				},
			},
		}).Do(context.Background())
	if err != nil {
		slog.Error("Error getting response:", err)
		return nil, err
	}

	var products []Product
	for _, hit := range res.Hits.Hits {
		hitByte, _ := hit.Source_.MarshalJSON()
		var product Product
		_ = json.Unmarshal(hitByte, &product)
		products = append(products, product)
	}

	return products, nil
}

func searchProductByName(name string) ([]Product, error) {
	res, err := esClient.Search().Index(productIndex).
		Request(&search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"name": {
						Query: name,
					},
				},
			},
		}).Do(context.Background())
	if err != nil {
		slog.Error("Error getting response:", err)
		return nil, err
	}

	var products []Product
	for _, hit := range res.Hits.Hits {
		hitByte, _ := hit.Source_.MarshalJSON()
		var product Product
		_ = json.Unmarshal(hitByte, &product)
		products = append(products, product)
	}

	return products, nil
}

func searchAndFilterProduct(name string, from, to int) ([]Product, error) {
	gte := types.Float64(from)
	lte := types.Float64(to)
	res, err := esClient.Search().Index(productIndex).
		Request(&search.Request{
			Query: &types.Query{
				Bool: &types.BoolQuery{
					Must: []types.Query{
						{
							Range: map[string]types.RangeQuery{
								"price": types.NumberRangeQuery{
									Gte: &gte,
									Lte: &lte,
								},
							},
						},
						{
							Match: map[string]types.MatchQuery{
								"name": {
									Query: name,
								},
							},
						},
					},
				},
			},
		}).Do(context.Background())
	if err != nil {
		slog.Error("Error getting response:", err)
		return nil, err
	}

	var products []Product
	for _, hit := range res.Hits.Hits {
		hitByte, _ := hit.Source_.MarshalJSON()
		var product Product
		_ = json.Unmarshal(hitByte, &product)
		products = append(products, product)
	}

	return products, nil
}

func searchProductWithHighlight(name string) ([]Product, error) {
	res, err := esClient.Search().Index(productIndex).
		Request(&search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"name": {
						Query: name,
					},
				},
			},
			Highlight: &types.Highlight{
				PreTags:  []string{"<b>"},
				PostTags: []string{"</b>"},
				Fields: map[string]types.HighlightField{
					"name": {},
				},
			},
		}).Do(context.Background())
	if err != nil {
		slog.Error("Error getting response:", err)
		return nil, err
	}

	var products []Product
	for _, hit := range res.Hits.Hits {
		hitByte, _ := hit.Source_.MarshalJSON()

		var product Product
		_ = json.Unmarshal(hitByte, &product)
		if highlightedName, ok := hit.Highlight["name"]; ok && len(highlightedName) == 1 {
			product.Name = highlightedName[0]
		}
		products = append(products, product)
	}

	return products, nil
}
