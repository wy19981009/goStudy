package main

import "fmt"

// 运算符
func operator() {
	fmt.Println("1+2=", 1+2)
	A := 20
	A++
	fmt.Println("A++=", A)

	fmt.Println("5>3的结果", 5 > 3)

	fmt.Println("!(4>3)的结果", !(4 > 3))

	fmt.Println("true && true的结果", true && true)

	fmt.Println("true && false的结果", true && false)

	a := 20
	fmt.Println("0<=a && a<=100结果", 0 <= a && a <= 100)
	fmt.Println("0<=a && a<=10结果", 0 <= a && a <= 10)
}
