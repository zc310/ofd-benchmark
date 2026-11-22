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

// Timeout reads the shared limit used by the benchmark drivers.
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

func Run(name string, args []string, outputFile string, timeout time.Duration) (time.Duration, time.Duration, int64, int64, bool) {
	fmt.Printf("=== %s ===\n", name)
	start := time.Now()
	// A failed or timed-out conversion must not leave a partial output for the
	// next invocation to mistake for a successfully created file.
	if err := os.Remove(outputFile); err != nil && !os.IsNotExist(err) {
		fmt.Printf("FAILED: remove previous output: %v\n", err)
		return time.Since(start), 0, 0, 0, false
	}

	duration, cpu, memory, err := runProcess(args, timeout)
	if err != nil {
		fmt.Printf("FAILED: %v\n", err)
		if err := os.Remove(outputFile); err != nil && !os.IsNotExist(err) {
			fmt.Printf("Failed to remove partial output: %v\n", err)
		}
		return duration, cpu, memory, 0, false
	}
	info, err := os.Stat(outputFile)
	if err != nil {
		fmt.Printf("FAILED: output file not created\n")
		return duration, cpu, memory, 0, false
	}
	fmt.Printf("Time: %v, CPU: %v, peak RSS: %d bytes\n", duration, cpu, memory)
	return duration, cpu, memory, info.Size(), true
}

// RunDirectory runs a converter writing multiple pages into a fresh directory.
// The caller owns the directory and validates the page set after execution.
func RunDirectory(name string, args []string, outputDir string, timeout time.Duration) (time.Duration, time.Duration, int64, error) {
	fmt.Printf("=== %s ===\n", name)
	if err := os.Mkdir(outputDir, 0o755); err != nil {
		return 0, 0, 0, err
	}
	return runProcess(args, timeout)
}

func runProcess(args []string, timeout time.Duration) (time.Duration, time.Duration, int64, error) {
	start := time.Now()
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
	if err := cmd.Start(); err != nil {
		return time.Since(start), 0, 0, err
	}
	stop := make(chan struct{})
	peak := make(chan int64, 1)
	go samplePeakRSS(cmd.Process.Pid, stop, peak)
	err := cmd.Wait()
	elapsed := time.Since(start)
	close(stop)
	memory := <-peak
	var cpu time.Duration
	if cmd.ProcessState != nil {
		cpu = cmd.ProcessState.UserTime() + cmd.ProcessState.SystemTime()
	}
	if ctx.Err() != nil {
		return elapsed, cpu, memory, fmt.Errorf("timed out after %s", timeout)
	}
	return elapsed, cpu, memory, err
}
