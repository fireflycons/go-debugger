// debugger/debugger_windows.go
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
	ret, _, _ := procIsDebuggerPresent.Call()
	return ret != 0
}
