package main

import "fmt"

type DB interface {
    Create() bool // creates a db file
    Add() bool // adds (string, string) pair to db
    Update() bool // updates an existing key
    Delete() bool // delets k/v for
    Get() string // returns value for key
}

func main() {
    fmt.Printf("Hello from gj\n");
}
