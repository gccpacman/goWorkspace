package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p1 = &Vertex{1, 2} // 类型为 *Vertex
)

func main() {
	v := Vertex{1, 2}
	p := &v //pointer
	q := v  //copy? 对其操作对v和q都没有影响
	//(*p).X = 1e9
	p.X = 1e9
	q.Y = 19
	fmt.Println(v, p, q)

	fmt.Println(v1, *p1, v2, v3)
}
