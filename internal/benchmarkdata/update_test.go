package benchmarkdata

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go.yaml.in/yaml/v4"
)

func TestUpdateCell(t *testing.T) {
	path := copyBenchmarkYAML(t)
	updates := []CellUpdate{{
		Report: "pdf-to-ofd", Table: "pdf-speed", Input: "hello.pdf",
		Converter: "go-zc310", Value: "42ms",
	}}
	if err := Update(path, updates, false); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), "values: [42ms, 111ms, 281ms") {
		t.Fatalf("updated value not found in YAML")
	}
}

func TestUpdateColumnWithReferenceCountInHeader(t *testing.T) {
	path := copyBenchmarkYAML(t)
	updates := []CellUpdate{{
		Report: "pdf-to-ofd", Table: "pdf-text-extraction", Input: "hello.pdf",
		Converter: "go-zc310", Value: "31",
	}}
	if err := Update(path, updates, false); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), `values: ["31", 10188, 9833`) {
		t.Fatal("updated header with reference count was not found")
	}
}

func TestResetKeepsReferenceBaseline(t *testing.T) {
	path := copyBenchmarkYAML(t)
	if err := Update(path, nil, true); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	text := string(raw)
	if !strings.Contains(text, "label: 数科版式阅读器, values: [1.00, 1.00, 1.00") {
		t.Fatal("reference baseline was reset")
	}
	if !strings.Contains(text, "values: [FAILED, FAILED, FAILED") {
		t.Fatal("benchmark cells were not reset")
	}
}

func TestWriteAndReadVersions(t *testing.T) {
	path := copyBenchmarkYAML(t)
	if err := WriteVersions(path, "v1-zc", "v1-ofdgo", "2.6.0"); err != nil {
		t.Fatal(err)
	}
	versions, err := ReadVersions(path)
	if err != nil {
		t.Fatal(err)
	}
	if versions[ZC310Module] != "v1-zc" || versions[OfdgoModule] != "v1-ofdgo" || versions[OfdrwArtifact] != "2.6.0" {
		t.Fatalf("unexpected versions: %#v", versions)
	}
}

func TestWriteEnvironment(t *testing.T) {
	path := copyBenchmarkYAML(t)
	values := map[string]string{
		EnvironmentOS:     "Ubuntu test",
		EnvironmentCPU:    "CPU test",
		EnvironmentCores:  "8",
		EnvironmentMemory: "16 GiB",
	}
	if err := WriteEnvironment(path, values); err != nil {
		t.Fatal(err)
	}
	var doc struct {
		Environment map[string]string `yaml:"benchmark_environment"`
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		t.Fatal(err)
	}
	for key, want := range values {
		if doc.Environment[key] != want {
			t.Fatalf("environment %s = %q, want %q", key, doc.Environment[key], want)
		}
	}
}

func copyBenchmarkYAML(t *testing.T) string {
	t.Helper()
	input, err := os.ReadFile("../../docs/benchmark-results.yaml")
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "benchmark-results.yaml")
	if err := os.WriteFile(path, input, 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}
