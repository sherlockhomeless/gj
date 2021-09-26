package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const (
	db_path = "/home/ml/go/src/github.com/sherlockhomeless/gj/db/db" //FIXME: Make flexible
	db_table = "paths"
)
var db_con *sql.DB

func init() {
	var err error
	log.Printf("Opening sqlite-db_ at %s\n", db_path)
	db_con, err = sql.Open("sqlite3", "db")
	if err != nil{
		log.Fatal(err)
	}
}

type PathDB struct {}

func (db *PathDB) Get (key string) []Entry {
	stmt, err := db_con.Prepare("SELECT * " + "FROM " + db_table + " where shorthand " )
	fmt.Println(stmt, err)
	return nil
}

func (db *PathDB) Add (key, value, extra string, prio int) bool{
	var stmt string
	

	return false
}



/**
  Add() bool // adds (string, string) pair to db_
  Update() bool // updates an existing key
  Delete() bool // delets k/v for
  Get() string // returns value for key
 */


