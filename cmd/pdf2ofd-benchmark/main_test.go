package main

import "testing"

func TestPDF2OFDPython(t *testing.T) {
	for _, path := range []string{"", "/tmp/pdf2ofd-venv/bin/python", "/tmp/venv with spaces/bin/python"} {
		t.Run(path, func(t *testing.T) {
			t.Setenv("PDF2OFD_PYTHON", path)
			want := path
			if want == "" {
				want = "python3"
			}
			if got := pdf2ofdPython(); got != want {
				t.Fatalf("pdf2ofdPython() = %q, want %q", got, want)
			}
		})
	}
}
