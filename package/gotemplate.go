package main

import (
	"./testpackage"
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
	testpackage.Func1()

	//testpackage.func2()
	//func2 is a private func
	//.\gotemplate.go:25: cannot refer to unexported name testpackage.func2
	//.\gotemplate.go:25: undefined: testpackage.func2

}

func (t T) Method1() {
	//..
}

func Func1() {
	//exported function Func1
	fmt.Println("Func1 is called")
}
