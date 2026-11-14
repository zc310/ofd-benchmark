// Package converter runs benchmark converters with a per-process time limit.
package converter

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

const DefaultTimeout = 15 * time.Minute

// Timeout reads the shared limit used by both benchmark drivers.
func Timeout() (time.Duration, error) {
	value := os.Getenv("BENCHMARK_CONVERTER_TIMEOUT")
	if value == "" {
		return DefaultTimeout, nil
	}
	timeout, err := time.ParseDuration(value)
	if err != nil || timeout <= 0 {
		return 0, fmt.Errorf("BENCHMARK_CONVERTER_TIMEOUT must be a positive duration, got %q", value)
	}
	return timeout, nil
}

func Run(name string, args []string, outputFile string, timeout time.Duration) (time.Duration, int64, bool) {
	fmt.Printf("=== %s ===\n", name)
	start := time.Now()
	// A failed or timed-out conversion must not leave a partial output for the
	// next invocation to mistake for a successfully created file.
	if err := os.Remove(outputFile); err != nil && !os.IsNotExist(err) {
		fmt.Printf("FAILED: remove previous output: %v\n", err)
		return time.Since(start), 0, false
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = os.TempDir()
	// Shell wrappers may spawn Java/Node processes. Kill their entire process
	// group on timeout rather than just the wrapper.
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Cancel = func() error {
		err := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		if errors.Is(err, syscall.ESRCH) {
			return os.ErrProcessDone
		}
		return err
	}
	cmd.WaitDelay = time.Second
	err := cmd.Run()
	duration := time.Since(start)
	if err != nil || ctx.Err() != nil {
		if ctx.Err() != nil {
			fmt.Printf("FAILED: timed out after %s\n", timeout)
		} else {
			fmt.Printf("FAILED: %v\n", err)
		}
		if err := os.Remove(outputFile); err != nil && !os.IsNotExist(err) {
			fmt.Printf("Failed to remove partial output: %v\n", err)
		}
		return duration, 0, false
	}
	info, err := os.Stat(outputFile)
	if err != nil {
		fmt.Printf("FAILED: output file not created\n")
		return duration, 0, false
	}
	fmt.Printf("Time: %v\n", duration)
	return duration, info.Size(), true
}
