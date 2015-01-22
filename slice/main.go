/*

数组
一个 slice 会指向一个序列的值，并且包含了长度信息。
[]T 是一个元素类型为 T 的 slice。

*/
package main

import "fmt"

func main() {

	AppendSlice()

	NilSlice()

	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	/*
		slice 可以重新切片
		创建一个新的 slice 值指向相同的数组。
	*/

	fmt.Println("p[1:4] ==", p[1:4])
	fmt.Println("p[:3] ==", p[:3]) // 省略下标代表从 0 开始
	fmt.Println("p[4:] ==", p[4:]) // 省略上标代表到 len(s) 结束

	/*
		slice 由函数 make 创建。这会分配一个零长度的数组并且返回一个 slice 指向这个数组：
		a := make([]int, 5)  // len(a)=5
		为了指定容量，可传递第三个参数到 `make`： 分配空间为0, 数组容量为5
		b := make([]int, 0, 5) // len(b)=0, cap(b)=5
		重新分配指针创建新的数组, 但是和b共享内存区
		b = b[:cap(b)] // len(b)=5, cap(b)=5
		b = b[1:]      // len(b)=4, cap(b)=4
	*/
	a := make([]int, 5)
	a[2] = 3
	printSlice("a", a)
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	//

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
