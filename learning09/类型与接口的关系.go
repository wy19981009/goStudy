package main

//
//import "fmt"
//
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
//type dog struct {
//	name string
//}
//
//type car struct {
//	brand string
//}
//
//// WashingMachine 洗衣机
//type WashingMachine interface {
//	wash()
//	dry()
//}
//
//// 甩干器
//type dryer struct{}
//
//// 海尔洗衣机
//type haier struct {
//	dryer //嵌入甩干器
//}
//
//// 实现Sayer接口
//func (d dog) say() {
//	fmt.Printf("%s会叫汪汪汪\n", d.name)
//}
//
//// 实现Mover接口
//func (d dog) move() {
//	fmt.Printf("%s会动\n", d.name)
//}
//
//// car类型实现Mover接口
//func (c car) move() {
//	fmt.Printf("%s速度70迈\n", c.brand)
//}
//
//// 实现WashingMachine接口的dry()方法
//func (d dryer) dry() {
//	fmt.Println("甩一甩")
//}
//
//// 实现WashingMachine接口的wash()方法
//func (h haier) wash() {
//	fmt.Println("洗刷刷")
//}
