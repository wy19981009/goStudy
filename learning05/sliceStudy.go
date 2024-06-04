package main

import "fmt"

func sliceStudy() {
	//var a = [3]int{1, 2, 3}
	////a[1:2] 生成了一个新的切片
	//fmt.Println(a, a[1:2])
	//
	//var highRiseBuilding [30]int
	//for i := 0; i < 30; i++ {
	//	highRiseBuilding[i] = i + 1
	//}
	//// 区间
	//fmt.Println(highRiseBuilding[10:15])
	//// 中间到尾部的所有元素
	//fmt.Println(highRiseBuilding[20:])
	//// 开头到中间指定位置的所有元素
	//fmt.Println(highRiseBuilding[:2])

	//// 声明字符串切片
	//var strList []string
	//// 声明整型切片
	//var numList []int
	//// 声明一个空切片
	//var numListEmpty = []int{}
	//// 输出3个切片
	//fmt.Println(strList, numList, numListEmpty)
	//// 输出3个切片大小
	//fmt.Println(len(strList), len(numList), len(numListEmpty))
	//// 切片判定空的结果
	//fmt.Println(strList == nil)
	//fmt.Println(numList == nil)
	//fmt.Println(numListEmpty == nil)
	//strList = append(strList, "ms的go教程")
	//strList = append(strList, "码神之路")
	//strList = append(strList, "ms的go教程")
	//fmt.Println(strList)

	//a := make([]int, 2)
	//b := make([]int, 2, 10)
	//fmt.Println(a, b)
	////容量不会影响当前的元素个数，因此 a 和 b 取 len 都是 2
	////但如果我们给a 追加一个 a的长度就会变为3
	//fmt.Println(len(a), len(b))
	//var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//myslice := numbers4[4:6]
	////这打印出来长度为2
	//fmt.Printf("myslice为 %d, 其长度为: %d\n", myslice, len(myslice))
	//myslice = myslice[:cap(myslice)]
	////为什么 myslice 的长度为2，却能访问到第四个元素
	//fmt.Printf("myslice的第四个元素为: %d", myslice[3])

	//slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := []int{5, 4, 3}
	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	//fmt.Println(slice2)
	//copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	//fmt.Println(slice1)

	// 设置元素数量为1000
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据 切片不会因为等号操作进行元素的复制
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素 引用数据的第一个元素将会发生变化
	fmt.Println(refData[0])
	// 打印复制切片的第一个和最后一个元素 由于数据是复制的，因此不会发生变化。
	fmt.Println(copyData[0], copyData[elementCount-1])
	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}
