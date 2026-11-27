package converter

import (
	"os"
	"path/filepath"
)

// OutputRoot is shared with the comparison and release scripts.
func OutputRoot() string {
	if root := os.Getenv("BENCHMARK_OUTPUT_DIR"); root != "" {
		if absolute, err := filepath.Abs(root); err == nil {
			return absolute
		}
		return root
	}
	return "/tmp/ofd-benchmark"
}
