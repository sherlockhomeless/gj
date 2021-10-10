package db

type DB interface {
	Add() bool // adds (string, string) pair to db_
	Update() bool // updates an existing key
	Delete() bool // delets k/v for
	Get() string // returns value for key
}

// Entry represents the different paths found in the history
type Entry struct {
	shorthand string // shorthand is the abbreviation used to jump, for example bin instead of ~/go/bin
	full_path string // full_path is the whole path on the system
	prio      int    // prio determines the order if more than one shorthand matches for a full_path
	extra     string // extra contains additional information that might be useful later
}

