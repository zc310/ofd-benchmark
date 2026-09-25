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

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	info, err := inputFile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	outputFile, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	if _, err := ofdgo.ConvertPDF(context.Background(), inputFile, info.Size(), outputFile, ofdgo.PDFImportOptions{}); err != nil {
		log.Fatal(err)
	}
}