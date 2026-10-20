package main

import (
	"context"
	"log"
	"os"

	"github.com/xiaoqidun/ofdgo"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: ofdgo-pdf2ofd <input.pdf> <output.ofd>")
	}

	input, err := ofdgo.OpenFile(context.Background(), os.Args[1], ofdgo.InputOptions{})
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	outputFile, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	if _, err := ofdgo.ConvertPDF(context.Background(), input, outputFile, ofdgo.PDFImportOptions{}); err != nil {
		log.Fatal(err)
	}
}