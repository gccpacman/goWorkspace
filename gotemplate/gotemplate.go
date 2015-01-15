package main

import (
	"./testa"
	"fmt"
)

const c = "C"

var v int = 5
var q int

type T struct{}

func init() {
	q = 7
}
func main() {
	var a int
	Func1()
	//...
	fmt.Println(a)
	fmt.Println(q)
	testa.Func1()
}

func (t T) Method1() {
	//..
}

func Func1() {
	//exported function Func1
	fmt.Println("Func1 is called")
}
