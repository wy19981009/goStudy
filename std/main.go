package main

import (
	"fmt"
	"io"
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
	fileReadDir()
}

type MyHandler struct {
}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "向浏览器发送数据")
}
