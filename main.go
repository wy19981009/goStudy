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

// 输入
func scanValue() {
	var a int
	fmt.Printf("请输入变量a\n")
	fmt.Scanf("%d", &a)
	fmt.Println("a=", a)
}

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

// 别名
func anotherName() {
	type bigint int64

	var a bigint

	fmt.Printf("a type is %T\n", a)

	type (
		long int64
		char byte
	)

	var ch char = 'b'
	var b long = 111
	fmt.Printf("ch=%c b=%d\n", ch, b)
}

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

func main() {
	// printf:字符串进行格式化
	// println: 回车换行
	// print 是否换行
	fmt.Println("Hello World")
	//changeValue()
	//scanValue()
	//convert()
	//anotherName()
	//operator()
}
