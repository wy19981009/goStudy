package main

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DB *sql.DB

func init() {
	// 执行main方法之前 会先执行init（初始化）方法
	db, err := sql.Open("mysql", "study:Yxl5201314...@tcp(gz-cdb-eg3lodhn.sql.tencentcdb.com:63499)/study?charset=utf8")
	if err != nil {
		log.Println("连接数据库异常")
		panic(err)
	}
	//最大空闲连接数，默认不配置，是2个最大空闲连接
	db.SetMaxIdleConns(5)
	//最大连接数，默认不配置，是不限制最大连接数
	db.SetMaxOpenConns(100)
	// 连接最大存活时间
	db.SetConnMaxLifetime(time.Minute * 3)
	//空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 1)
	err = db.Ping()
	if err != nil {
		log.Println("数据库连接失败")
		db.Close()
		panic(err)
	}
	DB = db
}

func save(username string, sex string, email string) {
	result, err := DB.Exec("insert into user (username, sex, email) values (?, ?, ?)", username, sex, email)
	if err != nil {
		log.Println("插入操作异常", err)
		return
	}
	id, _ := result.LastInsertId()
	log.Println("插入成功：", id)
}

type User struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

func query(id int) (*User, error) {
	rows, err := DB.Query("select * from user where user_id=? limit 1", id)
	if err != nil {
		log.Println("查询出现错误:", err)
		return nil, errors.New(err.Error())
	}
	user := new(User)
	for rows.Next() {
		if err := rows.Scan(&user.UserId, &user.Username, &user.Sex, &user.Email); err != nil {
			log.Println("scan error:", err)
			return nil, errors.New(err.Error())
		}
	}
	return user, nil
}

func main() {
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			return
		}
	}(DB)
	//save("wy666", "man", "test@qq.com")
	user, err := query(2)
	if err != nil {
		return
	}
	log.Println(user)
}
