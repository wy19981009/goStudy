package main

import "fmt"

func switchInitialize() {
	switch num := 4; num {
	case 1:
		fmt.Println("这个是数字1")
		break
	case 2:
		fmt.Println("这个是数字2")
		break
	case 3, 4, 5:
		fmt.Println("这个是数字345")
		break
	default:
		fmt.Println("这个是xxx")
		break

	}
}
