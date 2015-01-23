package main

import (
	"fmt"
	"math"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci is a function that returns
// a function that returns an int.
/*
练习：斐波纳契闭包
现在来通过函数做些有趣的事情。
实现一个 fibonacci 函数，返回一个函数（一个闭包）可以返回连续的斐波纳契数
*/
func fibonacci() func() int {
	v := 0
	p1 := 1
	p2 := 0
	return func() int {
		result := v
		v = p1 + p2
		p2 = p1
		p1 = v
		return result
	}
}

func main() {

	//练习
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	////////////////////////////

	//函数也是值
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))

	//函数闭包
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

}
