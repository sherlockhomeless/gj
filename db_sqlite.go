package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

/**
  Create() bool // creates a db file
  Add() bool // adds (string, string) pair to db
  Update() bool // updates an existing key
  Delete() bool // delets k/v for
  Get() string // returns value for key
 */

type DBSqlite struct {
	db_path string
	db_connection *sql.DB

}

func OpenDB (path string) *sql.DB{
	log.Println("Opening sqlite-db at %s", path)
	db, err := sql.Open("sqlite3,", "simple.sqlite")
	if err != nil{
		log.Fatal(err)
	}
	return db
}