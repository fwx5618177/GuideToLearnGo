package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    start := time.Now()
    for index, args := range os.Args[0:] {
        fmt.Println(index, args)
    }
    end := time.Now()

    fmt.Println("end:", start, end)
}