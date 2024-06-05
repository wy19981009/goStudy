package main

/*
* 数组是定长的，不可更改，在编译阶段就决定了
 */
func main() {
	//var arr [3]int
	//fmt.Println(arr[0])
	//fmt.Println(arr[1])
	//fmt.Println(arr[2])
	//var arr [3]int = [3]int{1, 2, 3}
	//var arr [3]int = [3]int{1, 2}
	//arr := [3]int{1, 2, 3}
	//arr := [...]int{1, 2, 3}
	//var arr [3]int
	//arr[0] = 5
	//arr[1] = 6
	//arr[2] = 7
	//for index, value := range arr {
	//	fmt.Printf("索引:%d,值：%d \n", index, value)
	//}
	//type arr3 [3]int
	////这样每次用arr3 代替[3]int，注意前面学过 定义一个类型后 arr3就是一个新的类型
	//var arr arr3
	//arr[0] = 2
	//for index, value := range arr {
	//	fmt.Printf("索引:%d,值：%d \n", index, value)
	//}
	//2 给索引为2的赋值 ，所以结果是 0,0,3
	//arr := [3]int{2: 3}
	//for index, value := range arr {
	//	fmt.Printf("索引:%d,值：%d \n", index, value)
	//}
	// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
	//var array [4][2]int
	//// 使用数组字面量来声明并初始化一个二维整型数组
	//array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	//// 声明并初始化数组中索引为 1 和 3 的元素
	//array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	//// 声明并初始化数组中指定的元素
	//array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	//for index, value := range array {
	//	fmt.Printf("索引:%d,值：%d \n", index, value)
	//}
	//sliceStudy()
	//mapStudy()
	nilStudy()
}
