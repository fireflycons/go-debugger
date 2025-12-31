//go:build linux

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fireflycons/go-debugger"
)

func main() {

	// Signal handler for CTRL-C etc.
	ctx, restoreSignals := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer restoreSignals()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	lastState := debugger.Attached()

	fmt.Printf("Debugger attached: %v. Try attaching a debugger. Process PID: %d\nPress CTRL-C to exit.\n", lastState, os.Getpid())

	debugger.Poll(ctx, time.Millisecond*500)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			state := debugger.Attached()
			if state != lastState {
				fmt.Printf("Debugger attached: %v\n", state)
				lastState = state
			}
		}
	}
}
