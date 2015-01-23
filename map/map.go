package main

import (
	"code.google.com/p/go-tour/wc"
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

/*
http://go-tour-zh.appspot.com/moretypes/19

练习：map
实现 `WordCount`。它应当返回一个含有 s 中每个 “词” 个数的 map。函数 wc.Test 针对这个函数执行一个测试用例，并输出成功还是失败。
你会发现 strings.Fields 很有帮助。
*/
func WordCount(s string) map[string]int {

	wordlist := strings.Fields(s)

	mapstr := make(map[string]int)

	for _, v := range wordlist {
		mapstr[v]++
	}

	//return map[string]int{"x": 1}
	return mapstr
}

func main() {

	wc.Test(WordCount)
	////////////////////

	//var m = map[string]Vertex{
	//	"Bell Labs": Vertex{
	//		40.68433, -74.39967,
	//	},
	//	"Google": Vertex{
	//		37.42202, -122.08408,
	//	},
	//}

	fmt.Println()
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)

	/**
		通过双赋值检测某个键存在：
	elem, ok = m[key]
	如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。
		**/

	key, exist := m["Bell Labs"]
	fmt.Println("The value:", key, "Present?", exist)

	//delete key
	delete(m, "Bell Labs")
	fmt.Println("The value:", m["Bell Labs"])

	v, ok := m["Bell Labs"]
	fmt.Println("The value:", v, "Present?", ok)

}
