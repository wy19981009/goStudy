package main

import (
	"fmt"
	"sync"
)

func mapStudy() {
	//var mapLit map[string]int
	//var mapAssigned map[string]int
	//mapLit = map[string]int{"one": 1, "two": 2}
	//mapAssigned = mapLit
	////mapAssigned 是 mapList 的引用，对 mapAssigned 的修改也会影响到 mapList 的值。
	//mapAssigned["two"] = 3
	//fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	//fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	//fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
	//
	//map2 := make(map[string]int, 100)
	//map2["cat"] = 66
	//map2["dog"] = 77
	//map2["pig"] = 963
	//for k, v := range map2 {
	//	fmt.Println(k, v)
	//}
	//fmt.Println("========================")
	//delete(map2, "cat")
	//for k, v := range map2 {
	//	fmt.Println(k, v)
	//}

	//// 创建一个int到int的映射
	//m := make(map[int]int)
	//// 开启一段并发代码
	//go func() {
	//	// 不停地对map进行写入
	//	for {
	//		m[1] = 1
	//	}
	//}()
	//// 开启一段并发代码
	//go func() {
	//	// 不停地对map进行读取
	//	for {
	//		_ = m[1]
	//	}
	//}()
	//// 无限循环, 让并发程序在后台执行
	//for {
	//}

	//sync.Map 不能使用 make 创建
	var scene sync.Map
	// 将键值对保存到sync.Map
	//sync.Map 将键和值以 interface{} 类型进行保存。
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	//遍历需要提供一个匿名函数，参数为 k、v，类型为 interface{}，每次 Range() 在遍历一个元素时，都会调用这个匿名函数把结果返回。
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
