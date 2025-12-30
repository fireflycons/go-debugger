//go:build linux

package debugger

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getParentPID(pid int) (int, error) {
	statPath := fmt.Sprintf("/proc/%d/stat", pid)
	data, err := os.ReadFile(statPath)
	if err != nil {
		return 0, err
	}

	fields := strings.Fields(string(data))
	if len(fields) < 4 {
		return 0, fmt.Errorf("unexpected /proc/[pid]/stat format")
	}

	return strconv.Atoi(fields[3])
}

func getProcessName(pid int) (string, error) {
	commPath := fmt.Sprintf("/proc/%d/comm", pid)
	data, err := os.ReadFile(commPath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}

func isBeingDebugged() bool {

	// Look up /proc/self/status
	// This will find any debugger attachment (inc strace etc)
	pid, err := tracerPID()

	if err == nil {
		return pid != 0
	}

	// Fall back to process tree
	pid = os.Getpid()

	for {
		ppid, err := getParentPID(pid)
		if err != nil || ppid == 0 {
			break
		}

		name, err := getProcessName(ppid)
		if err != nil {
			break
		}

		if name == "dlv" || strings.HasPrefix(name, "gdb") {
			return true
		}

		pid = ppid
	}

	return false
}

func poll(ctx context.Context, freq time.Duration) {

	go func() {

		ticker := time.NewTicker(freq)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if pid, err := tracerPID(); err == nil {
					attached.Store(pid != 0)
				}
			}
		}
	}()
}

func tracerPID() (int, error) {

	f, err := os.Open("/proc/self/status")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "TracerPid:") {
			// Split on any whitespace
			fields := strings.Fields(line)
			if len(fields) < 2 {
				return 0, nil
			}
			return strconv.Atoi(fields[1])
		}
	}

	return 0, scanner.Err()
}
