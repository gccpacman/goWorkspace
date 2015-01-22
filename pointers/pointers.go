package main

import "fmt"

func main() {

	i, j := 45, 270
	p := &i
	fmt.Println(*p)
	fmt.Println(p)

	*p = 21
	fmt.Println(*p)
	fmt.Println(p)

	p = &j
	fmt.Println(*p)
	fmt.Println(p)

	*p = *p / i
	fmt.Println(*p)
	fmt.Println(p)

	z := float64(j) / float64(i)
	q := &z
	//fmt.Println(z)

	fmt.Println(*q)
	fmt.Println(q)

	// cannot do that i don't know why
	// it seems Go cannot use one pointer to point both int and float32
	// that mean point also have type
	// it that right?
	// --- it is comfirmed.  see : http://www.golang-book.com/8/index.htm
	// pointers has type like *int  *float
	//p = &z
	//fmt.Println(*p)
	//fmt.Println(p)

}
