package main

import "fmt"

func main() {
	//fmt.Println("456")
	//x := 5
	//if x <= 0 {
	//	fmt.Println("为真进入这里")
	//	//go语言格式要求很严，else必须写在}后面
	//} else {
	//	fmt.Println("为假进入这里")
	//}
	//
	//var a int = 100
	//var b int = 200
	///* 判断条件 */
	//if a == 100 {
	//	/* if 条件语句为 true 执行 */
	//	if b == 200 {
	//		/* if 条件语句为 true 执行 */
	//		fmt.Printf("a 的值为 100 ， b 的值为 200\n")
	//	}
	//}
	//if a := 10; a > 5 {
	//	fmt.Println(a)
	//	return
	//}
	//sum := 0
	//for i := 0; i < 101; i++ {
	//	sum += i
	//
	//}
	//for {
	//	sum++
	//	if sum > 100 {
	//		//break是跳出循环
	//		break
	//	}
	//}
	//fmt.Println(sum)
	// 遍历, 决定处理第几行
	//for y := 1; y <= 9; y++ {
	//	// 遍历, 决定这一行有多少列
	//	for x := 1; x <= y; x++ {
	//		fmt.Printf("%d*%d=%d ", x, y, x*y)
	//	}
	//	// 手动生成回车
	//	fmt.Println()
	//}
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue OuterLoop
			}
		}
	}
}
