package main

import (
	"fmt"
	"time"
)

func main() {
	//for i := 0; i < 10; i++ {
	//	go hello(i)
	//}
	//fmt.Println("main goroutine done!")
	//time.Sleep(time.Second * 1)
	//go func(s string) {
	//	for i := 0; i < 2; i++ {
	//		fmt.Println(s)
	//	}
	//}("world")
	//// 主协程
	//for i := 0; i < 2; i++ {
	//	// 切一下，再次分配任务
	//	runtime.Gosched()
	//	fmt.Println("hello")
	//}
	//go func() {
	//	defer fmt.Println("A.defer")
	//	func() {
	//		defer fmt.Println("B.defer")
	//		// 结束协程
	//		runtime.Goexit()
	//		defer fmt.Println("C.defer")
	//		fmt.Println("B")
	//	}()
	//	fmt.Println("A")
	//}()
	//for {
	//}
	//runtime.GOMAXPROCS(2)
	//go a()
	//go b()
	//time.Sleep(time.Second)
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功")
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func hello(i int) {
	fmt.Println("Hello Goroutine", i)
}

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
	fmt.Println("A:", time.Now())
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
	fmt.Println("B:", time.Now())
}
