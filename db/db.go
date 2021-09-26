package db

type DB interface {
	Add() bool // adds (string, string) pair to db_
	Update() bool // updates an existing key
	Delete() bool // delets k/v for
	Get() string // returns value for key
}

type Entry struct {
	key string
	value string
	prio int
	extra string
}

