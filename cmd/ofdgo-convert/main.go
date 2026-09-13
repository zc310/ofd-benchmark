package main

import (
	"log"
	"os"

	"github.com/xiaoqidun/ofdgo"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: ofdgo-convert <file.ofd> [output.pdf]")
	}

	inputFile := os.Args[1]
	outputFile := "output_ofdgo.pdf"
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	}

	reader, err := ofdgo.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	pdfFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer pdfFile.Close()

	renderer := ofdgo.NewRenderer(reader, ofdgo.WithAnnotations(true))

	if err := renderer.RenderToMultiPagePDF(pdfFile); err != nil {
		log.Fatal(err)
	}
}
