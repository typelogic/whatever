package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        os.Exit(1)
    }
    inputj := os.Args[1]
    data,err := ioutil.ReadFile(inputj)
    if err == nil {
        fmt.Println(string(data))
        type Schema struct {
            Colors []string
        }

        var x Schema
        err = json.Unmarshal(data,&x)
        if err == nil {
            fmt.Println("OK")
            fmt.Println(x.Colors)
        } else {
            fmt.Println(err)
        }
    } else {
        fmt.Println(err)
    }
}
