package main

import (
	"fmt"
	"math"
)

//返回带一个int参数,返回一个int的函数
func adder(sum int) func(int) int {
	//sum := 0
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

	//函数也是值
	fmt.Println("函数也是值")
	hypot := func(x, y float64) float64 {
		fmt.Println("hypot() is called")
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("hypot=", hypot)
	fmt.Println("hypot(3, 4)=", hypot(3, 4))

	a := func() int {
		fmt.Println("a() is called")
		return 0
	}
	fmt.Println("a=", a)     //输出变量地址
	fmt.Println("a()=", a()) //执行函数, 输出结果

	xs := map[int]func() int{
		1: func() int {
			fmt.Println("xs[1] is called")
			return 1
		},
		2: func() int {
			fmt.Println("xs[2] is called")
			return 2
		},
	}

	fmt.Println("xs[1]=", xs[1])
	fmt.Println("xs[1]()=", xs[1]())
	fmt.Println("xs[2]=", xs[2])
	fmt.Println("xs[2]()=", xs[2]())

	//函数闭包
	fmt.Println("函数闭包")
	pos, neg := adder(0), adder(0)
	for i := 0; i < 10; i++ {
		fmt.Println(
			"pos(i)=", pos(i),
			"neg(-2*i)=", neg(-2*i),
			"adder(0)(i)=", adder(0)(i),
		)
	}

	//练习
	fmt.Println("递归:斐波纳契闭包")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println("loop", i, "--", f())
	}

}
