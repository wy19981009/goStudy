package main

func main() {
	//fmt.Println(maxInt(1, 10))
	//fmt.Println(maxInt(-1, -2))
	//fmt.Println(test(1, 2, "求和： %d"))
	//这是直接使用匿名函数
	//s1 := test(func() int { return 100 })
	//这是传入一个函数
	//s1 := test(fn)
	//fmt.Println(s1)
	//a, b, c := f1()
	//fmt.Println(a, b, c)
	//var a, b int = 1, 2
	/*
	   调用 swap() 函数
	   &a 指向 a 指针，a 变量的地址
	   &b 指向 b 指针，b 变量的地址
	*/
	//swap(&a, &b)
	//
	//fmt.Println(a, b)

	//s := []int{1, 2, 3}
	//res := test("sum: %d", s...) // slice... 展开slice
	//println(res)
	//这里将一个函数当做一个变量一样的操作。
	//getSqrt := func(a float64) float64 {
	//	return math.Sqrt(a)
	//}
	//fmt.Println(getSqrt(9))
	//func(data int) {
	//	fmt.Println("hello", data)
	//}(100) //(100)，表示对匿名函数进行调用，传递参数为 100。
	// 创建一个玩家生成器
	//generator := playerGen("ms")
	//// 返回玩家的名字和血量
	//name, hp := generator()
	//// 打印值
	//fmt.Println(name, hp)
	//
	//generator1 := playerGen1()
	//name1, hp1 := generator1("ms")
	//// 打印值
	//fmt.Println(name1, hp1)
	//var whatever = [5]int{1, 2, 3, 4, 5}
	//
	//for i := range whatever {
	//	defer fmt.Println(i)
	//}
	//start := time.Now()
	//log.Printf("开始时间为：%v", start)
	//defer func() {
	//	log.Printf("时间差：%v", time.Since(start))
	//}() // Now()此时已经copy进去了
	////不受这3秒睡眠的影响
	//time.Sleep(3 * time.Second)
	//
	//log.Printf("函数结束")
	//var whatever = [5]int{1, 2, 3, 4, 5}
	//for i, _ := range whatever {
	//	//函数正常执行,由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4.
	//	defer func() { fmt.Println(i) }()
	//}
	test()
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}
	}()

	panic("panic error!")
}

// 创建一个玩家生成器, 输入名称, 输出生成器
func playerGen(name string) func() (string, int) {
	// 血量一直为150
	hp := 150
	// 返回创建的闭包
	return func() (string, int) {
		// 将变量引用到闭包中
		return name, hp
	}
}

// 创建一个玩家生成器, 输入名称, 输出生成器
func playerGen1() func(string) (string, int) {
	// 血量一直为150
	hp := 150
	// 返回创建的闭包
	return func(name string) (string, int) {
		// 将变量引用到闭包中
		return name, hp
	}
}

//func test(s string, n ...int) string {
//	var x int
//	for _, i := range n {
//		x += i
//	}
//
//	return fmt.Sprintf(s, x)
//}

// 类型相同的相邻参数，参数类型可合并。
//func maxInt(n1, n2 int) int {
//	if n1 > n2 {
//		return n1
//	}
//	return n2
//}

//func test(x, y int, s string) (int, string) {
//	// 类型相同的相邻参数，参数类型可合并。 多返回值必须用括号。
//	n := x + y
//	return n, fmt.Sprintf(s, n)
//}

//func test(fn func() int) int {
//	return fn()
//}
//func fn() int {
//	return 200
//}
//
//// 支持返回值 命名 ，默认值为类型零值,命名返回参数可看做与形参类似的局部变量,由return隐式返回
//func f1() (names []string, m map[string]int, num int) {
//	m = make(map[string]int)
//	m["k1"] = 2
//
//	return
//}
//
///* 定义相互交换值的函数 */
//func swap(x, y *int) {
//	*x, *y = *y, *x
//}
