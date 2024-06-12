package main

import "fmt"

//type Point struct {
//	X int
//	Y int
//}
//
////type Player struct {
////	Name        string
////	HealthPoint int
////	MagicPoint  int
////}
//
////type Command struct {
////	Name    string // 指令名称
////	Var     *int   // 指令绑定的变量
////	Comment string // 指令的注释
////}
//
//func newCommand(name string, varRef *int, comment string) *Command {
//	return &Command{
//		Name:    name,
//		Var:     varRef,
//		Comment: comment,
//	}
//}
//
//var version = 1
//
//// 打印消息类型, 传入匿名结构体
//func printMsgType(msg *struct {
//	id   int
//	data string
//}) {
//	// 使用动词%T打印msg的类型
//	fmt.Printf("%T,\n msg:%v", msg, msg)
//}
//
//type Bag struct {
//	items []int
//}
//
//func (b *Bag) Insert(itemid int) {
//	b.items = append(b.items, itemid)
//}
//
//// Property 定义属性结构
//type Property struct {
//	value int // 属性值
//}
//
//// SetValue 设置属性值
//func (p *Property) SetValue(v int) {
//	// 修改p的成员变量
//	p.value = v
//}
//
//// Value 取属性值
//func (p *Property) Value() int {
//	return p.value
//}
//
//func main() {
//	//使用.来访问结构体的成员变量,结构体成员变量的赋值方法与普通变量一致。
//	//var p Point
//	//p.X = 1
//	//p.Y = 2
//	//var p = Point{
//	//	X: 1,
//	//	Y: 2,
//	//}
//	//var p = Point{
//	//	1,
//	//	2,
//	//}
//	//fmt.Printf("%v,x=%d,y=%d", p, p.X, p.Y)
//	//tank := new(Player)
//	//tank.Name = "ms"
//	//tank.HealthPoint = 300
//	//tank.MagicPoint = 100
//	//fmt.Println(tank)
//	//cmd := newCommand(
//	//	"version",
//	//	&version,
//	//	"show version",
//	//)
//	//fmt.Println(cmd)
//	// 实例化一个匿名结构体
//	//msg := &struct { // 定义部分
//	//	id   int
//	//	data string
//	//}{ // 值初始化部分
//	//	1024,
//	//	"hello",
//	//}
//	//printMsgType(msg)
//	//b := new(Bag)
//	//b.Insert(1001)
//	//fmt.Println(b.items)
//	// 实例化属性
//	p := new(Property)
//	// 设置值
//	p.SetValue(100)
//	// 打印值
//	fmt.Println(p.Value())
//}

//func main() {
//	// 实例化玩家对象，并设速度为0.5
//	p := NewPlayer(0.5)
//	// 让玩家移动到3,1点
//	p.MoveTo(Vec2{3, 1})
//	// 如果没有到达就一直循环
//	for !p.IsArrived() {
//		// 更新玩家位置
//		p.Update()
//		// 打印每次移动后的玩家位置
//		fmt.Println(p.Pos())
//	}
//	fmt.Printf("到达了：%v", p.Pos())
//}

//// 将int定义为MyInt类型
//type MyInt int
//
//// 为MyInt添加IsZero()方法
//func (m MyInt) IsZero() bool {
//	return m == 0
//}
//
//// 为MyInt添加Add()方法
//func (m MyInt) Add(other int) int {
//	return other + int(m)
//}
//
//func main() {
//	var b MyInt
//	fmt.Println(b.IsZero())
//	b = 1
//	fmt.Println(b.Add(2))
//	fmt.Println(b.IsZero())
//}

type User struct {
	id   int
	name string
}

type Manager struct {
	User
}

func (self *User) ToString() string { // receiver = &(Manager.User)
	return fmt.Sprintf("User: %p, %v", self, self)
}

func main() {
	m := Manager{User{1, "Tom"}}
	fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.ToString())
}
