//go:build !windows && !linux && !freebsd && !darwin

package debugger

import (
	"context"
	"time"
)

func isBeingDebugged() bool {
	// Not supported
	return false
}

func poll(context.Context, time.Duration) {
	// Not supported
}
