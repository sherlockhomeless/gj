package db

import (
	"testing"
        "time"
)

var test_path = "/tmp/gj_" + time.Now().GoString() + ".db_"

func setup_db() DBFile {
        Create(test_path)
        return OpenDB(test_path)
}

func TestDBFile_Add(t *testing.T) {
        db := setup_db()
}