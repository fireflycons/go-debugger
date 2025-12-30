//go:build !windows && !linux && !freebsd && !darwin

package debugger

func isBeingDebugged() bool {
	// Not supported
	return false
}
