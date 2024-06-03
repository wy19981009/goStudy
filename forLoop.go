package main

import "fmt"

func forLoop() {
	//sum := 0
	//for i := 0; i <= 100; i++ {
	//	sum = sum + i
	//}
	//fmt.Println(sum)
	str := "abc"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}

	for i, data := range str {
		fmt.Printf("str[%d]=%c\n", i, data)
	}
}
