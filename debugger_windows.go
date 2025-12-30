//go:build windows

package debugger

import (
	"syscall"
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
