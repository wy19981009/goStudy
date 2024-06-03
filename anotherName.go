package main

import "fmt"

// 别名
func anotherName() {
	type bigint int64

	var a bigint

	fmt.Printf("a type is %T\n", a)

	type (
		long int64
		char byte
	)

	var ch char = 'b'
	var b long = 111
	fmt.Printf("ch=%c b=%d\n", ch, b)
}
