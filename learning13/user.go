package main

import (
	"database/sql"
	"errors"
	"fmt"
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

func update(username string, id int) {
	ret, err := DB.Exec("update user set username=? where user_id=?", username, id)
	if err != nil {
		log.Println("更新出现问题:", err)
		return
	}
	affected, _ := ret.RowsAffected()
	fmt.Println("更新成功的行数:", affected)
}

func delete(id int) {
	ret, err := DB.Exec("delete from user where user_id=?", id)
	if err != nil {
		log.Println("删除出现问题:", err)
		return
	}
	affected, _ := ret.RowsAffected()
	fmt.Println("删除成功的行数:", affected)
}

func insertTx(username string) {
	tx, err := DB.Begin()
	if err != nil {
		log.Println("开启事务错误", err)
		return
	}
	ret, err := tx.Exec("insert into user (username,sex,email) values (?,?,?)", username, "man", "test@test.com")
	if err != nil {
		log.Println("事务sql执行出错:", err)
		return
	}
	id, _ := ret.LastInsertId()
	fmt.Println("插入成功:", id)
	if username == "lisi" {
		fmt.Println("回滚...")
		_ = tx.Rollback()
	} else {
		_ = tx.Commit()
	}
}

func main() {
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			return
		}
	}(DB)
	//save("wy666", "man", "test@qq.com")
	user, err := query(3)
	if err != nil {
		return
	}
	log.Println(user)
	//update("wy777", 2)
	//delete(2)
	//insertTx("wy008")
}
