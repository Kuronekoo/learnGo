package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("--------------testIf--------------")
	testMysql()

}

func testMysql() {
	dbUrl := "root:123456@tcp(127.0.0.1:3306)/party_website_db"
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	db.SetMaxOpenConns(100)
	if err != nil {
		panic(err)
	}
	qrySql := "select user_id,user_name from user_info_tab where user_id = ?"
	var user User
	err = db.QueryRow(qrySql, 1).Scan(&user.user_id, &user.user_name)

	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v", user)

	defer db.Close()
}

type User struct {
	user_id    uint64 `db:"pk"`
	user_name  string
	password   string
	email_addr string
	head_url   string
	ctime      uint32
	mtime      uint32
}
