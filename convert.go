package main

import "fmt"

// 类型转换
func convert() {
	var ch byte
	ch = 'A'
	var t int
	t = int(ch)
	fmt.Println("t=", t)

	var flag bool
	flag = true
	fmt.Printf("flag=%t\n\n", flag)
}
