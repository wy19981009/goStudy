package main

import (
	"fmt"
)

// DataWriter 定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
}

// 定义文件结构，用于实现DataWriter
type file struct {
}

// WriteData 实现DataWriter接口的WriteData方法
func (d *file) WriteData(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

func main() {
	// 实例化file
	//f := new(file)
	//// 声明一个DataWriter的接口
	//var writer DataWriter
	//// 将接口赋值f，也就是*file类型
	//writer = f
	//// 使用DataWriter接口进行数据写入
	//err := writer.WriteData("data")
	//if err != nil {
	//	return
	//}
	//var x Sayer
	//var y Mover
	//
	//var a = dog{name: "旺财"}
	//x = a
	//y = a
	//x.say()
	//y.move()
	//var x Mover
	//var a = dog{name: "旺财"}
	//var b = car{brand: "保时捷"}
	//x = a
	//x.move()
	//x = b
	//x.move()
	//var x animal
	//x = cat{name: "花花"}
	//x.move()
	//x.say()
	// 定义一个空接口x
	var x interface{}
	s := "ms的go教程"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}

//// Sayer 接口
//type Sayer interface {
//	say()
//}
//
//// Mover 接口
//type Mover interface {
//	move()
//}
//
//// 接口嵌套
//type animal interface {
//	Sayer
//	Mover
//}
//
//type cat struct {
//	name string
//}
//
//func (c cat) say() {
//	fmt.Println("喵喵喵")
//}
//
//func (c cat) move() {
//	fmt.Println("猫会动")
//}
