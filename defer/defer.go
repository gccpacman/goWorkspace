package main

import "fmt"

func main() {
	fmt.Println("hello")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	defer fmt.Println("where are you.")
	defer fmt.Println("world")

	fmt.Println("hello to you too")

	fmt.Println("see you soon")

}
