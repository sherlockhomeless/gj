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
// database connection that is opened by init-function
var db_con *sql.DB

// GetShort returns all entries where short == shorthand in db_con
func GetShort(short string) []Entry {
	rows, err := db_con.Query("SELECT * " + "FROM " + db_table + " where shorthand " )
	handle_err(err)
	defer rows.Close()

	results := make([]Entry, 5)
	for rows.Next() {
		var short, long, extra string
		var prio int
		err = rows.Scan(&short, &long, &prio, &extra)
		handle_err(err)
		results = append(results, Entry{short, long, prio, extra})
		}

	return results
}

func  Add (short, full, extra string, prio int) bool{
	query := fmt.Sprintf("INSERT INTO %s VALUES (%s, %s, %s, %s);", db_table, short, full, string(prio), extra)
	stmt, err := db_con.Prepare(query)


	if err != nil {
		panic(err)
	}

	return false
}



func init() {
	var err error
	log.Printf("Opening sqlite-db_ at %s\n", db_path)
	db_con, err = sql.Open("sqlite3", "db")
	if err != nil{
		log.Fatal(err)
	}
}

func handle_err(err error){
	if (err != nil){
		panic(err)
	}
}