package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on \n")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
		fallthrough
	case "linux":
		fmt.Println("Linux.")
		fallthrough
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
