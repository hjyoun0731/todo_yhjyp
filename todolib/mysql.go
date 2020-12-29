package todolib

import (
	"database/sql"
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

	_, err = db.Exec("DROP TABLE IF EXISTS version")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db drop complete")

	_, err = db.Exec("CREATE TABLE version (id serial PRIMARY KEY, version VARCHAR(10), name VARCHAR(10), updateTime INTEGER)")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db table create complete")
}

// DbVerInsert put data to db
func DbVerInsert(arg version) {
	db, err := sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/tododb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println(arg)
	result, err := db.Exec("INSERT INTO version(version, name, updateTime) VALUES (?, ?, ?)",arg.Version, arg.Name, arg.UpdateTime)
	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected()
	if n == 1 {
		MakeLog("1 row insert success")
	}
}

// DbAcctInsert put data to db
func DbAcctInsert(arg account) {
	db, err := sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/tododb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO account(userid, password) VALUES (?, ?)",arg.UserId, arg.Password)
	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected()
	if n == 1 {
		MakeLog("1 row insert success")
	}
}

// DbQuery get data from db
func DbQuery(col string,table string) string {
	db, err := sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/tododb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var getData string
	err = db.QueryRow("SELECT " + col + " FROM "+ table +" ORDER BY id DESC").Scan(&getData)
	if err != nil {
		log.Fatal(err)
	}
	return getData
}

// DbDelete delete data except newest version
func DbDelete(arg string){
	db, err := sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/tododb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM version WHERE version != ?", arg)
	if err != nil {
		log.Fatal(err)
	}
	MakeLog("DbDelete success")
}