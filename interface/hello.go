package main

import "fmt"

type Circle struct {
	x, y, r float64
}

type Bird struct {
	name string
}

type IFly interface {
	Fly()
}

func (b *Bird) Fly() {
	fmt.Println("Fly as Bird, " + b.name + "!")
}

func main() {

	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c)

	f := Bird{"Anna"}
	fmt.Println(f.name)

	//var fly IFly = new(Bird)
	//var fly IFly = IFly(new(Bird))
	var fly IFly = IFly(&Bird{"John"})
	fly.Fly()
	fmt.Printf("Hello, World!\n")

}
