package db

import (
	"fmt"

	"goTest/base"

	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	userName  string = "root"
	password  string = "123456"
	ipAddrees string = "127.0.0.1"
	port      int    = 3306
	dbName    string = "testdb"
	charset   string = "utf8"
)

func DbConnection() *sqlx.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	Db, err := sqlx.Open("mysql", dns)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}

	return Db
}

func QueryAllStudent(db *sqlx.DB) {
	rows, err := db.Query("select sid, name, age, sex, school,class from student")
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return

	}
	fmt.Println("Sid, Name, Age, Sex, School,Class")
	for rows.Next() {
		var student base.Student
		rows.Scan(&student.Sid, &student.Name, &student.Age, &student.Sex, &student.School, &student.Class)

		res, _ := json.Marshal(student)
		fmt.Println(string(res))
	}
	rows.Close()
}

func InsertNewStudent(db *sqlx.DB, student base.Student) {
	stm, err := db.Prepare("INSERT INTO student set name=?,age=?,class=?,school=?,sex=?")
	if err != nil {
		return
	}

	stm.Exec(student.Name, student.Age, student.Class, student.School, student.Sex)
}

func DelStudent(db *sqlx.DB, student base.Student) {
	stm, err := db.Prepare("delete from student where name =?")
	if err != nil {
		return
	}
	stm.Exec(student.Name)
}
