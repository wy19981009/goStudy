package main

import "fmt"

// 输入
func scanValue() {
	var a int
	fmt.Printf("请输入变量a\n")
	fmt.Scanf("%d", &a)
	fmt.Println("a=", a)
}
