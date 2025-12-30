//go:build freebsd && !386

package debugger

import (
	"syscall"

	"golang.org/x/sys/unix"
)

// From <sys/ptrace.h> on FreeBSD
const ptDenyAttach = 31

func isBeingDebugged() bool {
	// ptrace(PT_DENY_ATTACH, 0, NULL, 0)
	_, _, errno := syscall.Syscall6(
		unix.SYS_PTRACE,
		uintptr(ptDenyAttach),
		0,
		0,
		0,
		0,
		0,
	)

	// If already being traced, kernel returns EPERM
	return errno == unix.EPERM
}
