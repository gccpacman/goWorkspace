package main

import "fmt"

func NilSlice() {
	//nil slice 空数组
	var z []int
	//z := make([]int, 0, 0)   //不是nil, 而是len=0, cap=0的slice
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}
