package todolib

import (
	"database/sql"
	"fmt"
	"log"
)

// DbInit start db
func DbInit() {
	db, err := sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/tododb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("db open complete")

	_, err = db.Exec("DROP TABLE IF EXISTS test")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db drop complete")

	_, err = db.Exec("CREATE TABLE test (id serial PRIMARY KEY, version VARCHAR(10), name VARCHAR(10), updateTime INTEGER)")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db table create complete")
}

// DbInsert put data to db
func DbInsert() {
	db, err := sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/tododb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO test(version,name,updateTime) VALUES ('0.0.1','yhjyp',0)")
	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
}

// DbQuery get data from db
func DbQuery(arg string) string {
	db, err := sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/tododb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var getData string
	err = db.QueryRow("SELECT " + arg + " FROM test WHERE id = 1").Scan(&getData)
	if err != nil {
		log.Fatal(err)
	}
	return getData
}
