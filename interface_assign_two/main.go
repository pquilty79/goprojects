package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
)

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {             
        fmt.Println(scanner.Text())  


    }
}