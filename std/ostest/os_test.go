package ostest

import (
	"fmt"
	"os"
	"testing"
)

// 创建文件
func TestCreate(t *testing.T) {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f: %v\n", f)
	}
}

func TestMkdir(t *testing.T) {
	err := os.MkdirAll("msa/b/c", os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func TestRemoveFile(t *testing.T) {
	err := os.Remove("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
func TestRemoveDir(t *testing.T) {
	err := os.Remove("ms")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func TestRenameFile(t *testing.T) {
	err := os.Rename("test.txt", "test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
func TestRenameDir(t *testing.T) {
	err := os.Rename("msa", "ms")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
