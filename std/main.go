package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func fileStat() {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("file info : %#v", fileInfo)
}

func fileRead() {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var body []byte
	for {
		buf := make([]byte, 4)
		n, err := f.Read(buf)
		if err == io.EOF {
			//读完了
			break
		}
		fmt.Printf("读到的位置:%d \n", n)
		body = append(body, buf[:n]...)
	}
	fmt.Printf("内容：%s \n", body)
}

func fileReadAt() {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := make([]byte, 5)
	n, err := f.ReadAt(buf, 6)
	fmt.Printf("内容：%s \n", buf[:n])
}

func fileReadDir() {
	f, err := os.Open("ostest")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//-1代表所有
	dirs, err := f.ReadDir(-1)
	if err != nil {
		panic(err)
	}
	for _, v := range dirs {
		fmt.Println("is dir:", v.IsDir())
		fmt.Println("dir name :", v.Name())
	}
}

var logger *log.Logger

func init() {
	logFile, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	logger = log.New(logFile, "[Mylog]", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	//http.ListenAndServe(":8088", &MyHandler{})
	//空格 换行输入都可以
	//var (
	//	name    string
	//	age     int
	//	married bool
	//)
	////fmt.Scan(&name, &age, &married)
	//fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	//var (
	//	name    string
	//	age     int
	//	married bool
	//)
	//r := strings.NewReader("10 false 张三")
	//
	//n, err := fmt.Fscanf(r, "%d %t %s", &age, &married, &name)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Fscanf:%v\n", err)
	//}
	//
	//fmt.Println(name, age, married)
	//
	//fmt.Println(n)
	//fileStat()
	//fileRead()
	//fileReadAt()
	//fileReadDir()
	//var wg sync.WaitGroup
	//wg.Add(2)
	////NewTimer 创建一个 Timer，它会在最少过去时间段 d 后到期，向其自身的 C 字段发送当时的时间
	//timer1 := time.NewTimer(2 * time.Second)
	//
	////NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调
	////整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可
	////以释放相关资源。
	//ticker1 := time.NewTicker(2 * time.Second)
	//
	//go func(t *time.Ticker) {
	//	defer wg.Done()
	//	for {
	//		<-t.C
	//		fmt.Println("get ticker1", time.Now().Format("2006-01-02 15:04:05"))
	//	}
	//}(ticker1)
	//
	//go func(t *time.Timer) {
	//	defer wg.Done()
	//	for {
	//		<-t.C
	//		fmt.Println("get timer", time.Now().Format("2006-01-02 15:04:05"))
	//		//Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。如果调用时t
	//		//还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
	//		t.Reset(2 * time.Second)
	//	}
	//}(timer1)
	//
	//wg.Wait()
	//log.Print("this is a log")
	//log.Printf("this is a log: %d", 100) // 格式化输出
	//name := "zhangsan"
	//age := 20
	//log.Println(name, " ", age)
	//defer fmt.Println("发生了 panic错误！")
	//log.Print("this is a log")
	//log.Panic("this is a panic log ")
	//fmt.Println("运行结束。。。")
	//defer fmt.Println("defer。。。")
	//log.Print("this is a log")
	//log.Fatal("this is a fatal log")
	//fmt.Println("运行结束。。。")
	//i := log.Flags()
	//fmt.Printf("i: %v\n", i)
	//log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	//log.Print("this is a log")
	//s := log.Prefix()
	//fmt.Printf("s: %v\n", s)
	//log.SetPrefix("[MyLog] ")
	//s = log.Prefix()
	//fmt.Printf("s: %v\n", s)
	//log.Print("this is a log...")
	//f, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//if err != nil {
	//	log.Panic("打开日志文件异常")
	//}
	//log.SetOutput(f)
	//log.Print("this is a file log...")
	logger.Println("自定义logger")
}

type MyHandler struct {
}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "向浏览器发送数据")
}
