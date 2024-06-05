package main

import "fmt"

func main() {
	//fmt.Println(maxInt(1, 10))
	//fmt.Println(maxInt(-1, -2))
	//fmt.Println(test(1, 2, "求和： %d"))
	//这是直接使用匿名函数
	//s1 := test(func() int { return 100 })
	//这是传入一个函数
	s1 := test(fn)
	fmt.Println(s1)
}

// 类型相同的相邻参数，参数类型可合并。
//func maxInt(n1, n2 int) int {
//	if n1 > n2 {
//		return n1
//	}
//	return n2
//}

//func test(x, y int, s string) (int, string) {
//	// 类型相同的相邻参数，参数类型可合并。 多返回值必须用括号。
//	n := x + y
//	return n, fmt.Sprintf(s, n)
//}

func test(fn func() int) int {
	return fn()
}
func fn() int {
	return 200
}
