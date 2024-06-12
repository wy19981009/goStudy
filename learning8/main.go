package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

type Command struct {
	Name    string // 指令名称
	Var     *int   // 指令绑定的变量
	Comment string // 指令的注释
}

func newCommand(name string, varRef *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     varRef,
		Comment: comment,
	}
}

var version = 1

// 打印消息类型, 传入匿名结构体
func printMsgType(msg *struct {
	id   int
	data string
}) {
	// 使用动词%T打印msg的类型
	fmt.Printf("%T\n, msg:%v", msg, msg)
}

func main() {
	//使用.来访问结构体的成员变量,结构体成员变量的赋值方法与普通变量一致。
	//var p Point
	//p.X = 1
	//p.Y = 2
	//var p = Point{
	//	X: 1,
	//	Y: 2,
	//}
	//var p = Point{
	//	1,
	//	2,
	//}
	//fmt.Printf("%v,x=%d,y=%d", p, p.X, p.Y)
	//tank := new(Player)
	//tank.Name = "ms"
	//tank.HealthPoint = 300
	//tank.MagicPoint = 100
	//fmt.Println(tank)
	//cmd := newCommand(
	//	"version",
	//	&version,
	//	"show version",
	//)
	//fmt.Println(cmd)
	// 实例化一个匿名结构体
	msg := &struct { // 定义部分
		id   int
		data string
	}{ // 值初始化部分
		1024,
		"hello",
	}
	printMsgType(msg)
}
