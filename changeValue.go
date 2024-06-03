package main

import "fmt"

// 打印变量
func changeValue() {
	a := 10
	b := "abc"
	c := 'a'
	d := 3.14

	fmt.Printf("%T %T %T %T\n", a, b, c, d)

	// %d: 整数 %s: 字符串 %c: 单个字符 %f: 浮点数
	fmt.Printf("%d %s %c %f\n", a, b, c, d)
}
