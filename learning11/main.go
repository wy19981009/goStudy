package main

import (
	"fmt"
	"sync"
	"sync/atomic"
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
	//ch := make(chan int)
	//go recv(ch)
	//ch <- 15
	//fmt.Println("发送成功")
	//ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	//ch <- 10
	//fmt.Println("发送成功")
	//c := make(chan int)
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		c <- i
	//	}
	//	close(c)
	//}()
	//for {
	//	if data, ok := <-c; ok {
	//		fmt.Println(data)
	//	} else {
	//		break
	//	}
	//}
	//fmt.Println("main结束")
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//// 开启goroutine将0~100的数发送到ch1中
	//go func() {
	//	for i := 0; i < 100; i++ {
	//		ch1 <- i
	//	}
	//	close(ch1)
	//}()
	//// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	//go func() {
	//	for {
	//		i, ok := <-ch1 // 通道关闭后再取值ok=false
	//		if !ok {
	//			break
	//		}
	//		ch2 <- i * i
	//	}
	//	close(ch2)
	//}()
	//// 在主goroutine中从ch2中接收值打印
	//for i := range ch2 { // 通道关闭后会退出for range循环
	//	fmt.Println(i)
	//}
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//go counter(ch1)
	//go squarer(ch2, ch1)
	//printer(ch2)
	// 2个管道
	//output1 := make(chan string)
	//output2 := make(chan string)
	//// 跑2个子协程，写数据
	//go test1(output1)
	//go test2(output2)
	//// 用select监控
	//select {
	//case s1 := <-output1:
	//	fmt.Println("s1=", s1)
	//case s2 := <-output2:
	//	fmt.Println("s2=", s2)
	//}

	// 创建2个管道
	//intChan := make(chan int, 1)
	//stringChan := make(chan string, 1)
	//go func() {
	//	//time.Sleep(2 * time.Second)
	//	intChan <- 1
	//}()
	//go func() {
	//	stringChan <- "hello"
	//}()
	//select {
	//case value := <-intChan:
	//	fmt.Println("int:", value)
	//case value := <-stringChan:
	//	fmt.Println("string:", value)
	//}
	//fmt.Println("main结束")
	// 创建管道
	//output1 := make(chan string, 10)
	//// 子协程写数据
	//go write(output1)
	//// 取数据
	//for s := range output1 {
	//	fmt.Println("res:", s)
	//	time.Sleep(time.Second)
	//}

	//wg.Add(2)
	//go add()
	//go add()
	//wg.Wait()
	//fmt.Println(x)

	//start := time.Now()
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go write()
	//}
	//
	//for i := 0; i < 1000; i++ {
	//	wg.Add(1)
	//	go read()
	//}
	//
	//wg.Wait()
	//end := time.Now()
	//fmt.Println(end.Sub(start))
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		// go add()       // 普通版add函数 不是并发安全的
		// go mutexAdd()  // 加锁版add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}

var x int64
var l sync.Mutex
var wg sync.WaitGroup

// 普通版加函数
func add() {
	// x = x + 1
	x++ // 等价于上面的操作
	wg.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

//
//var (
//	x      int64
//	wg     sync.WaitGroup
//	lock   sync.Mutex
//	rwlock sync.RWMutex
//)
//
//func write() {
//	lock.Lock() // 加互斥锁
//	// rwlock.Lock() // 加写锁
//	x = x + 1
//	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
//	//rwlock.Unlock()                   // 解写锁
//	lock.Unlock() // 解互斥锁
//	wg.Done()
//}
//
//func read() {
//	lock.Lock() // 加互斥锁
//	//rwlock.RLock()               // 加读锁
//	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
//	//rwlock.RUnlock()             // 解读锁
//	lock.Unlock() // 解互斥锁
//	wg.Done()
//}

//var x int64
//var wg sync.WaitGroup
//var lock sync.Mutex

//func add() {
//
//	for i := 0; i < 5000; i++ {
//		lock.Lock()
//		x = x + 1
//		lock.Unlock()
//	}
//	wg.Done()
//}

//func write(ch chan string) {
//	for {
//		select {
//		// 写数据
//		case ch <- "hello":
//			fmt.Println("write hello")
//		default:
//			fmt.Println("channel full")
//		}
//		time.Sleep(time.Millisecond * 500)
//	}
//}

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
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
