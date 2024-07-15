package main

import (
	"fmt"
	"io"
	"log"
	"math"
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

type Person struct {
	Name    string
	Age     int
	Email   string
	Married bool
	Parent  []string
}

type NewInts []uint

func (n NewInts) Len() int {
	return len(n)
}

func (n NewInts) Less(i, j int) bool {
	fmt.Println(i, j, n[i] < n[j], n)
	return n[i] < n[j]
}

func (n NewInts) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
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
	//logger.Println("自定义logger")

	//var b = []byte("seafood") //强制类型转换
	//
	//a := bytes.ToUpper(b)
	//fmt.Println(a, b)
	//
	//c := b[0:4]
	//c[0] = 'A'
	//fmt.Println(c, b)

	//s1 := "Φφϕ kKK"
	//s2 := "ϕΦφ KkK"
	//
	//// 看看 s1 里面是什么
	//for _, c := range s1 {
	//	fmt.Printf("%-5x", c)
	//}
	//fmt.Println()
	//// 看看 s2 里面是什么
	//for _, c := range s2 {
	//	fmt.Printf("%-5x", c)
	//}
	//fmt.Println()
	//// 看看 s1 和 s2 是否相似
	//fmt.Println(bytes.EqualFold([]byte(s1), []byte(s2)))
	//bs := [][]byte{ //[][]byte 字节切片 二维数组
	//	[]byte("Hello World !"),
	//	[]byte("Hello 世界！"),
	//	[]byte("hello golang ."),
	//}
	//
	//f := func(r rune) bool {
	//	return bytes.ContainsRune([]byte("!！. "), r) //判断r字符是否包含在"!！. "内
	//}
	//
	//for _, b := range bs { //range bs  取得下标和[]byte
	//	fmt.Printf("去掉两边: %q\n", bytes.TrimFunc(b, f)) //去掉两边满足函数的字符
	//}
	//
	//for _, b := range bs {
	//	fmt.Printf("去掉前缀: %q\n", bytes.TrimPrefix(b, []byte("Hello "))) //去掉前缀
	//}

	//b := []byte("  Hello   World !  ")
	//fmt.Printf("b: %q\n", b)
	//fmt.Printf("%q\n", bytes.Split(b, []byte{' '}))
	//
	//fmt.Printf("%q\n", bytes.Fields(b))
	//
	//f := func(r rune) bool {
	//	return bytes.ContainsRune([]byte(" !"), r)
	//}
	//fmt.Printf("%q\n", bytes.FieldsFunc(b, f))
	//b := []byte("hello golang") //字符串强转为byte切片
	//sublice1 := []byte("hello")
	//sublice2 := []byte("Hello")
	//fmt.Println(bytes.Contains(b, sublice1)) //true
	//fmt.Println(bytes.Contains(b, sublice2)) //false
	//
	//s := []byte("hellooooooooo")
	//sep1 := []byte("h")
	//sep2 := []byte("l")
	//sep3 := []byte("o")
	//fmt.Println(bytes.Count(s, sep1)) //1
	//fmt.Println(bytes.Count(s, sep2)) //2
	//fmt.Println(bytes.Count(s, sep3)) //9
	//s := []byte("hello,world")
	//old := []byte("o")
	//news := []byte("ee")
	//fmt.Println(string(bytes.Replace(s, old, news, 0)))  //hello,world
	//fmt.Println(string(bytes.Replace(s, old, news, 1)))  //hellee,world
	//fmt.Println(string(bytes.Replace(s, old, news, 2)))  //hellee,weerld
	//fmt.Println(string(bytes.Replace(s, old, news, -1))) //hellee,weerld
	//
	//s1 := []byte("你好世界")
	//r := bytes.Runes(s1)
	//fmt.Println("转换前字符串的长度：", len(s1)) //12
	//fmt.Println("转换后字符串的长度：", len(r))  //4

	//s1 := []int{1, 2, 3}
	//i := append(s1, 4)
	//fmt.Printf("i: %v\n", i)
	//
	//s2 := []int{7, 8, 9}
	//i2 := append(s1, s2...)
	//fmt.Printf("i2: %v\n", i2)

	//p := Person{
	//	Name:    "zhangsan",
	//	Age:     20,
	//	Email:   "zhangsan@mail.com",
	//	Married: true,
	//}
	//b, _ := json.Marshal(p)
	//jsonstr := string(b)
	//fmt.Printf("b: %v\n", jsonstr)
	//pe := &Person{}
	//json.Unmarshal([]byte(jsonstr), pe)
	//fmt.Printf("person: %v", pe)
	//f, _ := os.Open("test.json")
	//defer f.Close()
	//
	//d := json.NewDecoder(f)
	//var v map[string]any
	//d.Decode(&v)
	//
	//fmt.Printf("v: %v\n", v)
	//p := Person{
	//	Name:   "zhangsan",
	//	Age:    20,
	//	Email:  "zhangsan@mail.com",
	//	Parent: []string{"Daddy", "Mom"},
	//}
	//f, _ := os.OpenFile("test1.json", os.O_WRONLY|os.O_CREATE, 077)
	//defer f.Close()
	//
	//d := json.NewEncoder(f)
	//d.Encode(p)
	//n := []uint{1, 3, 2}
	//sort.Sort(NewInts(n))
	//fmt.Println(n)
	fmt.Println(math.IsNaN(12321.321321))
	fmt.Println(math.Ceil(1.13456))
	fmt.Println(math.Floor(2.9999))
	fmt.Println(math.Trunc(2.9999))
	fmt.Println(math.Abs(2.999312323132141665374)) //2.9993123231321417
	fmt.Println(math.Abs(2.999312323132141465374)) //2.9993123231321412
	fmt.Println(math.Max(1000, 200))
	fmt.Println(math.Min(1000, 200))
	fmt.Println(math.Mod(123, 0))  //NaN
	fmt.Println(math.Mod(123, 10)) //3
}

type MyHandler struct {
}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "向浏览器发送数据")
}
