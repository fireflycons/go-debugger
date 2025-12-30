package main

import (
	"fmt"
	"os"

	"github.com/fireflycons/go-debugger"
)

func main() {

	if debugger.Attached() {
		fmt.Println("Debugger is attached")
		b, _ := os.ReadFile("/proc/self/status")
		fmt.Println(string(b))
	} else {
		fmt.Println("Debugger is not attached")
	}
}
