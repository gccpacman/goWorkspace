package main

import "fmt"

func mydefer() int {
	defer func(x int) {
		fmt.Println("defer test", x+3)
	}(5)
	return 0
}

func f() (rec int) {
	//rec := 10
	defer func(x int) {
		rec = rec + 10 //可以改变return的值
	}(10)
	return 5
}

func main() {

	defer func() {
		fmt.Println("defer test")
	}()

	fmt.Println(f())
	fmt.Println(mydefer())
	//demo
	fmt.Println("hello")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	defer fmt.Println("where are you.")
	defer fmt.Println("world")

	fmt.Println("hello to you too")

	fmt.Println("see you soon")

}
