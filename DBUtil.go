package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type article struct {
	title string
	body  string
}

var db *sql.DB

func initDB() {
	var err error

	db, err = sql.Open("mysql", "root:@/godb")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func insert(a *article) {
	initDB()

	stmt, err := db.Prepare("insert into article(title, content) values(?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	stmt.Exec(a.title, a.body)

	fmt.Println(a.title)

	defer db.Close()
}
