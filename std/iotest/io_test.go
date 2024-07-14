package iotest

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	f, err := os.Open("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 12) // 实例化一个长度为4的[]byte
	var body []byte
	f.Seek(1, 0)
	for {
		n, err2 := f.Read(buf) // 将内容读至buf
		if n == 0 || err2 == io.EOF {
			fmt.Println("文件以读取完毕")
			break
		}
		body = append(body, buf[:n]...)

	}
	fmt.Println(string(body))
}

func TestWrite(t *testing.T) {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0775) // 以读写模式打开文件，并且在写操作时将数据附加到文件尾部
	f.Write([]byte(" hello golang"))
	f.Close()
}
