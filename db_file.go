package main

import (
    "bufio"
    "log"
    "os"
    "strings"
)

/**
    * Uses a simple text file for persisting data
    * Format KEY:VALUE
 */

type DBFile struct {
    path string
    entries map[string]string
}

func OpenDB(db_path string) DBFile{
    db := DBFile{path: db_path}

    f, err := os.Open(db.path)

    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    key_value := []string{}
    for scanner.Scan() {
        key_value = strings.Split(scanner.Text(), ":")
        db.entries[key_value[0]] = key_value[1]
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return db
}

func Create (path string){
    empty_file, err := os.Create(path)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("created empty file at %s", path)
    empty_file.Close()
}

func (db DBFile) Add(path string){

}
func (db DBFile) Update(path string){

}
func (db DBFile) Delete(path string){

}
func (db DBFile) Get(path string){

}
