// Package debugger contains utilities for managing debug sessions of the application the package is imported into.
//
// Supported archtectures:
//   - Linux
//   - Windows
//   - MacOS
//   - FreeBSD (64 bit)
package debugger

// Attached is true if this process was launched by a debugger.
//
// This variable is initialized by an init() function, therefore there
// is zero cost to use it. Example use case would be to dynamically set
// a context timeout such that it does not cancel in the middle of a
// debugging session.
//
// Note that the value will not be true if you attach a debugger to
// an already running process.
var Attached bool

func init() {
	Attached = isBeingDebugged()
}
