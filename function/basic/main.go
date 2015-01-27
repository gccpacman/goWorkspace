package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) retFloat() float64 {
	return -float64(f)
}

func (f MyFloat) toInt() int {
	return int(f)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	p := Vertex{7, 7}
	fmt.Println(p.Abs())
	/*
	   你可以对包中的 任意 类型定义任意方法，而不仅仅是针对Struct。
	   但是，不能对来自其他包的类型或基础类型定义方法。
	*/
	fv := MyFloat(-math.Sqrt2)
	fmt.Println(fv.retFloat())
	fmt.Println(fv.toInt())

	fp := MyFloat(-0.153322)
	fmt.Println(fp.retFloat())
	fmt.Println(fp.toInt())
}
