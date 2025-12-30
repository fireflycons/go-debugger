//go:build linux

package debugger

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
	pid := os.Getpid()

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
