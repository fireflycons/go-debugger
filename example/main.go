package main

import (
	"fmt"

	"github.com/fireflycons/go-debugger"
)

func main() {

	if debugger.Attached {
		fmt.Println("Debugger is attached")
	} else {
		fmt.Println("Debugger is not attached")
	}
}
