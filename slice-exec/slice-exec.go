/**
http://go-tour-zh.appspot.com/moretypes/14
http://tour.golang.org/moretypes/14

not very sure how to do it.
**/

package main

import (
	//"code.google.com/p/go-tour/pic"
	"fmt"
)

func main() {

	slice := make([]int, 6)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	var array [100]int
	// IF[ var array [100]string ] IS WRONG.  -------\slice-exec.go:23: cannot use array[0:99] (type []string) as type []int in assignment
	slice = array[0:99]
	slice[98] = 'a'
	fmt.Println(slice)
	//slice[99] = 'a'
}
