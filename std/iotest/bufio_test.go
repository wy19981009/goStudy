package iotest

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestNewReader(t *testing.T) {
	//s := strings.NewReader("ABCDEFG")
	//str := strings.NewReader("12345")
	//br := bufio.NewReader(s)
	//b, _ := br.ReadString('\n')
	//fmt.Println(b)
	//br.Reset(str)
	//b, _ = br.ReadString('\n')
	//fmt.Println(b)

	//s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	//br := bufio.NewReader(s)
	//p := make([]byte, 10)
	//
	//for {
	//	n, err := br.Read(p)
	//	if err == io.EOF {
	//		break
	//	} else {
	//		fmt.Printf("string(p[0:n]): %v\n", string(p[0:n]))
	//	}
	//}
	//s := strings.NewReader("ABCDEFG")
	//br := bufio.NewReader(s)
	//
	//c, _ := br.ReadByte()
	//fmt.Printf("%c\n", c)
	//
	//c, _ = br.ReadByte()
	//fmt.Printf("%c\n", c)
	//
	//br.UnreadByte()
	//c, _ = br.ReadByte()
	//fmt.Printf("%c\n", c)

	s := strings.NewReader("ABC\nDEF\r\nGHI\r\nGHI")
	br := bufio.NewReader(s)

	w, isPrefix, _ := br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)
}
