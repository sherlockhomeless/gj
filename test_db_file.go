package main

import (
        "testing"
        "time"
)

var test_path = "/tmp/" + time.Now().GoString()

func Test_OpenDB(t testing.T){
        Create(test_path)
}