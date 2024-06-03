package main

import "fmt"

func useOfIf() {
	var a int
	fmt.Printf("请输入数字a：\n")
	fmt.Scan(&a)
	if a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a != 10")
	}

	// if语句初始化，然后在赋值 和程序上面的a不一样
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a != 10")
	}

	fmt.Println("a=", a)

	a = 8
	if a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else {
		fmt.Println("这是不可能的")
	}

	b := 10
	if b == 10 {
		fmt.Println("b == 10")
	}
}
