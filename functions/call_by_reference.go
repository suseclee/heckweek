package main

import "flag"
import "fmt"

type arrayFlags []string



func (i *arrayFlags) String() string {
    return "my string representation"
}

func (i *arrayFlags) Set(value string) (err error) {
    *i = append(*i, value)
    return nil
}

var myFlags arrayFlags



func main() {
    flag.Var(&myFlags, "list1", "Some description for this param.")
    flag.Parse()
    fmt.Println(myFlags)
    fmt.Println(myFlags.String())
    myFlags.Set("sss")
    fmt.Println(myFlags)
}