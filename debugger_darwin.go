//go:build darwin

package debugger

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func isBeingDebugged() bool {
	// No easy way to walk the process heirarchy with process names
	// other than to parse the output of ps.
	cmd := exec.Command("ps", "-o", "pid,ppid,comm", "-ax")
	out, err := cmd.Output()
	if err != nil {
		return false
	}

	pid := os.Getpid()
	parents := map[int]int{}
	names := map[int]string{}

	scanner := bufio.NewScanner(bytes.NewReader(out))
	scanner.Scan() // skip header

	for scanner.Scan() {
		var p, pp int
		var name string
		fmt.Sscanf(scanner.Text(), "%d %d %s", &p, &pp, &name)
		parents[p] = pp
		names[p] = name
	}

	for pid != 0 {
		if names[pid] == "dlv" || strings.HasPrefix(names[pid], "gdb") || strings.Contains(names[pid], "/dlv") || strings.Contains(names[pid], "/gdb") {
			return true
		}
		pid = parents[pid]
	}

	return false
}

func poll(context.Context, time.Duration) {
	// Not supported
}
