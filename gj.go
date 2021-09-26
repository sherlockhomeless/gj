package main

import (
    "fmt"
    "log"
)



func main() {
    set_logging()

    fmt.Printf("Hello from gj\n");
}


func set_logging(){
    log.SetFlags(log.Lshortfile)
}
