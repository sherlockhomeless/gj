package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type db struct {
    Path string
}

func AddPair(p string, sh string, path string) error {
    d, err := loadDB(p)
    if err != nil{
        panic("cant load db")
        return err
    }
    d[sh] = path
    return nil
}

func DeletePair(p string, sh string) error {
    
}
func UpdatePair(p string, sh string, path string) error {

}

func loadDB(p string) (map[string]string, error){
    file, err := os.Open(p)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    var d map[string]string
    for sc.Scan(){
        line := sc.Text()
        pair := strings.Split(line, ":")
        key, value := pair[0], pair[1]
        d[key] = value
    }
    return d, nil
}