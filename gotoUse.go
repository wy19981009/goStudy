package main

import "fmt"

func gotoUse() {
	//	fmt.Println("111111")
	//	goto END //goto是关键字 END是用户起的名字，它叫标签，goto不能跨函数调用
	//	fmt.Println("222222")
	//END:
	//	fmt.Println("3333333")

	for i := 0; i < 5; i++ {
		for {
			fmt.Println(i)
			goto LABLE
		}
		fmt.Println("this is test")
	}
LABLE:
	fmt.Println("it is over")
}
