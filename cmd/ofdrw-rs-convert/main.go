package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/lihongjie0209/ofdrw-rs/bindings/go/ofdrw"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: ofdrw-rs-convert <file.ofd> [output.pdf]")
	}

	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	if os.Getenv("OFDRW_LIB_PATH") == "" {
		libName := map[string]string{
			"windows": "ofdrw_ffi.dll",
			"darwin":  "libofdrw_ffi.dylib",
			"linux":   "libofdrw_ffi.so",
		}[runtime.GOOS]
		if libName == "" {
			log.Fatalf("unsupported operating system: %s", runtime.GOOS)
		}
		_ = os.Setenv("OFDRW_LIB_PATH", filepath.Join(filepath.Dir(exePath), "ofdrw-rs", libName))
	}

	outputFile := "output_ofdrw_rs.pdf"
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	}
	doc, err := ofdrw.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer doc.Close()

	pdf, err := doc.ToPDF()
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(outputFile, pdf, 0644); err != nil {
		log.Fatal(err)
	}
}
