package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("%s\n", runtime.Version())
	fmt.Println(runtime.GOROOT)
	//fmt.Println(runtime.GOOS)

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
