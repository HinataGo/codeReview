package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	id   int
	name string
)

func main() {

	db, err := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/go_learn")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	rows, err := db.Query("select * from go_learn.demo01 ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
