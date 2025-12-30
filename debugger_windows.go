//go:build windows

package debugger

import (
	"context"
	"syscall"
	"time"
)

var (
	kernel32              = syscall.NewLazyDLL("kernel32.dll")
	procIsDebuggerPresent = kernel32.NewProc("IsDebuggerPresent")
)

func isBeingDebugged() bool {
	// Windows makes it easy to know if the current process is being debugged.
	ret, _, _ := procIsDebuggerPresent.Call()
	return ret != 0
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
				attached.Store(isBeingDebugged())
			}
		}
	}()
}
