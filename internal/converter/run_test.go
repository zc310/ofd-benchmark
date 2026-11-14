package converter

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	for _, tc := range []struct {
		value string
		want  time.Duration
		valid bool
	}{
		{"", DefaultTimeout, true},
		{"10m", 10 * time.Minute, true},
		{"250ms", 250 * time.Millisecond, true},
		{"0s", 0, false},
		{"-1s", 0, false},
		{"invalid", 0, false},
	} {
		t.Run(tc.value, func(t *testing.T) {
			t.Setenv("BENCHMARK_CONVERTER_TIMEOUT", tc.value)
			got, err := Timeout()
			if (err == nil) != tc.valid || got != tc.want {
				t.Fatalf("Timeout() = %s, %v; want %s, valid=%v", got, err, tc.want, tc.valid)
			}
		})
	}
}

func TestRun(t *testing.T) {
	for _, tc := range []struct {
		name, script string
		ok           bool
	}{
		{"success", `printf new > "$1"`, true},
		{"failure", `printf partial > "$1"; exit 1`, false},
		{"missing output", `exit 0`, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			output := filepath.Join(t.TempDir(), "output.pdf")
			if err := os.WriteFile(output, []byte("stale"), 0o644); err != nil {
				t.Fatal(err)
			}
			_, size, ok := Run(tc.name, []string{"/bin/sh", "-c", tc.script, "converter", output}, output, 5*time.Second)
			if ok != tc.ok {
				t.Fatalf("Run() success = %v, want %v", ok, tc.ok)
			}
			if ok && size != 3 {
				t.Fatalf("size = %d, want 3", size)
			}
			if !ok {
				if _, err := os.Stat(output); !os.IsNotExist(err) {
					t.Fatalf("failed conversion left an output: %v", err)
				}
			}
		})
	}
}

func TestRunTimeoutKillsChildrenAndContinues(t *testing.T) {
	dir := t.TempDir()
	output := filepath.Join(dir, "output.pdf")
	marker := filepath.Join(dir, "child-finished")
	// The child would recreate the partial output if only the wrapper were killed.
	script := `printf partial > "$1"; (sleep 1; printf survived > "$1"; printf done > "$2") & wait`
	duration, _, ok := Run("timeout", []string{"/bin/sh", "-c", script, "converter", output, marker}, output, 250*time.Millisecond)
	if ok || duration > 3*time.Second {
		t.Fatalf("timed-out converter: success=%v, duration=%s", ok, duration)
	}
	_, _, ok = Run("next converter", []string{"/bin/sh", "-c", `printf new > "$1"`, "converter", output}, output, 5*time.Second)
	if !ok {
		t.Fatal("converter after timeout did not succeed")
	}
	time.Sleep(1100 * time.Millisecond)
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Fatalf("child survived process group timeout: %v", err)
	}
	contents, err := os.ReadFile(output)
	if err != nil || string(contents) != "new" {
		t.Fatalf("next converter output = %q, %v", contents, err)
	}
}
