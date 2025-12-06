// Package exit does, in the spirit of Advent of Code, provide quick and dirty ways to fail.
package exit

import (
	"fmt"
	"os"
	"runtime"
)

// If prints err and exits if err != nil.
func If(err error) {
	_, _, line, _ := runtime.Caller(1)
	if err != nil {
		fmt.Printf("Error:%d: %v", line, err)
		os.Exit(1)
	}
}

// PanicIf panics if err != nil.
func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
