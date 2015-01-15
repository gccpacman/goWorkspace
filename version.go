package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("%s\n", runtime.Version())
	fmt.Println(runtime.GOROOT)
	fmt.Println(runtime.GOOS)
}
