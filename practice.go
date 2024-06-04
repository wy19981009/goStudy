package main

import "fmt"

func practice() {
	s1 := "localhost:8080"
	b := []byte(s1)
	b[len(b)-1] = '1'
	s2 := string(b)
	fmt.Println(s2)
}
