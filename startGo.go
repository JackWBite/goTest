package main

import (
	"fmt"
	"goTest/base"

	"goTest/db"
)

func main() {
	fmt.Printf("hello, Go!")
	base.SayHi()

	tom := base.Student{base.Human{"tom", 1}, 0, 10, "一年级一班", "光明路一小"}

	var dbConn = db.DbConnection()
	defer dbConn.Close()

	db.InsertNewStudent(dbConn, tom)

	db.QueryAllStudent(dbConn)

	db.DelStudent(dbConn, tom)

	jim := base.Human{"jim", 1}

	/*** 多线程 ****/
	go tom.SayHi()

	go func() {
		fmt.Println("匿名函数代码线程执行")
	}()

	jim.SayHi()
	/***多线程****/
}
