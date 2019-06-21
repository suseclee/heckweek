package main

import (
    "log"
    "os"
    "strings"
    "flag"
    "fmt"
    "bufio"
)

func main() {

    path := flag.String("path", "journalctl.log" , "log file path")
    module := flag.String("module", "kernel", "module you want to debug")

    flag.Parse()
    f, err := os.Open(*path)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    r := bufio.NewReader(f)
    for{
        s, err := r.ReadString('\n')
        if err != nil {
            break
        }
        if strings.Contains(s, *module) {
            fmt.Print(s)
        }
    }
}