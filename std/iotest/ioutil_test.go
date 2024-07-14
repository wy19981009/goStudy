package iotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestNopCloser(t *testing.T) {
	f, err := os.Open("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	readCloser := ioutil.NopCloser(f)
	fmt.Printf("readCloser: %v\n", readCloser)
}

func TestReadAll(t *testing.T) {
	f, _ := os.Open("a.txt") // File实现了Reader
	defer f.Close()

	b, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("string(b): %v\n", string(b))
}

func TestWriteAll(t *testing.T) {
	ioutil.WriteFile("a.txt", []byte("hello world"), 0664)
}

func TestTempDir(t *testing.T) {
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("dir: %v\n", dir)
	// defer os.RemoveAll(dir) // 销毁临时目录

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}
}
