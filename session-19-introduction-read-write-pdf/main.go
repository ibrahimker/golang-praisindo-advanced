package main

import (
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
)

const (
	basedir      = "session-19-introduction-read-write-pdf"
	htmlTemplate = "static/template.html"
	pdfOutput    = "static/output.pdf"
)

func main() {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	// Create a new input page from an URL
	page := wkhtmltopdf.NewPage(basedir + "/" + htmlTemplate)
	// Add to document
	pdfg.AddPage(page)
	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}
	// Write buffer contents to file on disk
	err = pdfg.WriteFile(basedir + "/" + pdfOutput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done
}
