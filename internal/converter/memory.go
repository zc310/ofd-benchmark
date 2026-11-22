package converter

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const memorySampleInterval = 10 * time.Millisecond

// RSS is sampled for the live process tree, not summed individual peak RSS.
// A zero result means no resident memory was observed.
func samplePeakRSS(root int, stop <-chan struct{}, result chan<- int64) {
	ticker := time.NewTicker(memorySampleInterval)
	defer ticker.Stop()
	var peak int64
	for {
		if rss := processTreeRSS("/proc", root, int64(os.Getpagesize())); rss > peak {
			peak = rss
		}
		select {
		case <-stop:
			result <- peak
			return
		case <-ticker.C:
		}
	}
}

func processTreeRSS(proc string, root int, pageSize int64) int64 {
	pending := []int{root}
	seen := make(map[int]bool)
	var total int64
	for len(pending) > 0 {
		pid := pending[len(pending)-1]
		pending = pending[:len(pending)-1]
		if seen[pid] {
			continue
		}
		seen[pid] = true
		base := filepath.Join(proc, strconv.Itoa(pid))
		if data, err := os.ReadFile(filepath.Join(base, "statm")); err == nil {
			fields := strings.Fields(string(data))
			if len(fields) >= 2 {
				if pages, err := strconv.ParseInt(fields[1], 10, 64); err == nil && pages > 0 {
					total += pages * pageSize
				}
			}
		}
		// Children can be created by any thread, not only the thread-group leader.
		tasks, err := os.ReadDir(filepath.Join(base, "task"))
		if err != nil {
			continue
		}
		for _, task := range tasks {
			data, err := os.ReadFile(filepath.Join(base, "task", task.Name(), "children"))
			if err != nil {
				continue
			}
			for _, child := range strings.Fields(string(data)) {
				if pid, err := strconv.Atoi(child); err == nil && pid > 0 {
					pending = append(pending, pid)
				}
			}
		}
	}
	return total
}

// FormatPeakRSS uses decimal MB, matching the report generator's size units.
func FormatPeakRSS(bytes int64) string {
	if bytes <= 0 {
		return "FAILED"
	}
	return fmt.Sprintf("%.2fMB", float64(bytes)/1_000_000)
}
