package main

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
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
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		slog.Error("Error creating es client:", err)
		os.Exit(1)
	}

	res, err := es.Search().
		Index("users").
		Request(&search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"name": {Query: "ibam"},
				},
			},
		}).Do(context.Background())
	if err != nil {
		slog.Error("Error getting response:", err)
		os.Exit(1)
	}
	slog.Info("es response", slog.Any("es res", res.Hits))
	for _, hit := range res.Hits.Hits {
		slog.Info("data", slog.Any("hit", hit.Source_))
	}

	res, err = es.Search().
		Index("users").
		Request(&search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"name": {Query: "budi"},
				},
			},
		}).Do(context.Background())
	if err != nil {
		slog.Error("Error getting response:", err)
		os.Exit(1)
	}
	for _, hit := range res.Hits.Hits {
		slog.Info("data", slog.Any("hit", hit.Source_))
		hitByte, _ := hit.Source_.MarshalJSON()
		var u User
		_ = json.Unmarshal(hitByte, &u)
		slog.Info("converted data to struct", slog.Any("user", u))
	}
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
