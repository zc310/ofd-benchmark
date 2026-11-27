package converter

import (
	"path/filepath"
	"testing"
)

func TestOutputRoot(t *testing.T) {
	t.Setenv("BENCHMARK_OUTPUT_DIR", "")
	if got := OutputRoot(); got != "/tmp/ofd-benchmark" {
		t.Fatalf("default output root: %q", got)
	}
	t.Setenv("BENCHMARK_OUTPUT_DIR", "relative-results")
	want, err := filepath.Abs("relative-results")
	if err != nil {
		t.Fatal(err)
	}
	if got := OutputRoot(); got != want {
		t.Fatalf("output root = %q, want %q", got, want)
	}
}
