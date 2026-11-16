package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zc310/ofd-benchmark/internal/benchmarkdata"
)

func main() {
	path := flag.String("yaml", "docs/benchmark-results.yaml", "benchmark results YAML")
	zc310 := flag.String("zc310", "", "tested github.com/zc310/ofd version")
	ofdgo := flag.String("ofdgo", "", "tested github.com/xiaoqidun/ofdgo version")
	ofdrw := flag.String("ofdrw", "", "tested org.ofdrw:ofdrw-converter version")
	osName := flag.String("os", "", "tested operating system")
	cpu := flag.String("cpu", "", "tested CPU model and architecture")
	cores := flag.String("cores", "", "number of logical CPU cores")
	memory := flag.String("memory", "", "available memory")
	check := flag.Bool("check", false, "print true if any version differs from the recorded versions")
	write := flag.Bool("write", false, "record the tested versions in YAML")
	flag.Parse()

	if *zc310 == "" || *ofdgo == "" || *ofdrw == "" {
		fatal(fmt.Errorf("-zc310, -ofdgo, and -ofdrw are required"))
	}
	if *check == *write {
		fatal(fmt.Errorf("choose exactly one of -check or -write"))
	}
	if *check {
		versions, err := benchmarkdata.ReadVersions(*path)
		if err != nil {
			fatal(err)
		}
		changed := versions[benchmarkdata.ZC310Module] != *zc310 ||
			versions[benchmarkdata.OfdgoModule] != *ofdgo ||
			versions[benchmarkdata.OfdrwArtifact] != *ofdrw
		if changed {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
		return
	}
	if err := benchmarkdata.WriteVersions(*path, *zc310, *ofdgo, *ofdrw); err != nil {
		fatal(err)
	}
	if *osName != "" || *cpu != "" || *cores != "" || *memory != "" {
		if *osName == "" || *cpu == "" || *cores == "" || *memory == "" {
			fatal(fmt.Errorf("when recording environment, -os, -cpu, -cores, and -memory are all required"))
		}
		environment := map[string]string{
			benchmarkdata.EnvironmentOS:     *osName,
			benchmarkdata.EnvironmentCPU:    *cpu,
			benchmarkdata.EnvironmentCores:  *cores,
			benchmarkdata.EnvironmentMemory: *memory,
		}
		if err := benchmarkdata.WriteEnvironment(*path, environment); err != nil {
			fatal(err)
		}
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "benchmark-versions:", err)
	os.Exit(1)
}
