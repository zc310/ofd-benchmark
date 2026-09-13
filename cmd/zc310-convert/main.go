package main

import (
	"log"
	"os"

	"github.com/zc310/ofd/pkg/converter"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: zc310-convert <file.ofd> [output.pdf]")
	}

	inputFile := os.Args[1]
	outputFile := "output_zc310.pdf"
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	}

	pdfFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer pdfFile.Close()

	if err := converter.PDF(inputFile, pdfFile); err != nil {
		log.Fatal(err)
	}
}
