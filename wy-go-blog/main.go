package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "wy.com"
	indexData.Desc = "wy.com is a good blog"
	t := template.New("index.html")
	// 拿到当前的路径
	path, _ := os.Getwd()
	files, _ := t.ParseFiles(path + "/template/index.html")
	files.Execute(w, indexData)
}

func main() {
	// 程序入口，一个项目 只能有一个 main.go 入口
	// web程序
	service := http.Server{
		Addr: "127.0.0.1:7888",
	}
	http.HandleFunc("/", index)
	if err := service.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
