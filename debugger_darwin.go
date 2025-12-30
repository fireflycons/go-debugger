// debugger/debugger_darwin.go
//go:build darwin

package debugger

import "golang.org/x/sys/unix"

// From sys/proc.h on Darwin
const pTraced = 0x000008

func isBeingDebugged() bool {
	info, err := unix.SysctlKinfoProc("kern.proc.pid", unix.Getpid())
	if err != nil {
		return false
	}

	return info.Proc.P_flag&pTraced != 0
}
