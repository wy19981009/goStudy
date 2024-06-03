package main

import "fmt"

func useOfSwitch() {
	b := 10
	switch b {
	case 1:
		fmt.Println("这个是数字1")
		break
	case 2:
		fmt.Println("这个是数字2")
		break
	case 3:
		fmt.Println("这个是数字3")
		break
	case 4:
		fmt.Println("这个是数字4")
		break
	default:
		fmt.Println("这个是xxx")
		break
	}
}
