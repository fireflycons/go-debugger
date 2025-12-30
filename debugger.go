// Package debugger contains utilities for managing debug sessions of the application the package is imported into.
//
// Supported archtectures:
//   - Linux
//   - Windows
//   - MacOS
//   - FreeBSD (64 bit)
package debugger

// Attached is true if this process was launched by a debugger.
var Attached = isBeingDebugged()
