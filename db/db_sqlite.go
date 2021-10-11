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
	paths_schemata = "(shorthand, fullpath, priority, extra)"
)
// database connection that is opened by init-function
var db_con *sql.DB

// GetShort returns all entries where shorthand == shorthand in db_con
func GetShort(short string) []Entry {
	query := fmt.Sprintf("SELECT * FROM %s where shorthand = \"%s\";", db_table, short)
	rows, err := db_con.Query(query)
	log.Printf("GetShort: Query: %s", query)

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

// AddEntry creates an Entry and adds it to the db_con
func AddEntry(e *Entry) {
	query := fmt.Sprintf("INSERT INTO %s%s VALUES(\"%s\", \"%s\", %s, \"%s\");", db_table, paths_schemata, e.shorthand, e.full_path, fmt.Sprint(e.prio), e.extra)
	log.Printf("AddEntry: Query: %s", query)
	_, err := db_con.Exec(query)
	handle_err(err)
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