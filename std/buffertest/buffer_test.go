package buffertest

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRW(t *testing.T) {
	rd := bytes.NewBufferString("Hello World!")
	b := rd.Bytes()
	buf := make([]byte, 6)
	rd.Read(buf)
	fmt.Printf("%s\n\n", b)
	fmt.Printf("%s\n", rd.String())

	rd.Read(buf)
	fmt.Printf("%s\n", rd.String())
	fmt.Printf("%s\n", b)

	rd.Reset()
	fmt.Printf("%s\n", rd.String())
}

func TestTW2(t *testing.T) {
	data := "123456789"
	//通过[]byte创建Reader
	re := bytes.NewReader([]byte(data))
	//返回未读取部分的长度
	fmt.Println("re len : ", re.Len())
	//返回底层数据总长度
	fmt.Println("re size : ", re.Size())

	fmt.Println("---------------")

	buf := make([]byte, 2)
	for {
		//读取数据
		n, err := re.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

func TestTW3(t *testing.T) {
	data := "123456789"
	//通过[]byte创建Reader
	re := bytes.NewReader([]byte(data))

	buf := make([]byte, 2)

	re.Seek(0, 0)
	//设置偏移量
	for {
		//一个字节一个字节的读
		b, err := re.ReadByte()
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}
	fmt.Println("----------------")

	re.Seek(0, 0)
	off := int64(0)
	for {
		//指定偏移量读取
		n, err := re.ReadAt(buf, off)
		if err != nil {
			break
		}
		off += int64(n)
		fmt.Println(off, string(buf[:n]))
	}
}
