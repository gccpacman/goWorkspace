package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("%s\n", runtime.Version())
	fmt.Println(runtime.GOROOT)
	fmt.Println(runtime.GOOS)
	fmt.Println(time.Now())
}
