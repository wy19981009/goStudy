package main

import "fmt"

// NewInt 将NewInt定义为int类型
type NewInt int

// IntAlias 将int取一个别名叫IntAlias
type IntAlias = int

func otherName() {
	// 将a声明为NewInt类型
	var a NewInt
	// 查看a的类型名 main.NewInt
	fmt.Printf("a type: %T\n", a)
	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名 int
	//IntAlias 类型只会在代码中存在，编译完成时，不会有 IntAlias 类型。
	fmt.Printf("a2 type: %T\n", a2)
}
