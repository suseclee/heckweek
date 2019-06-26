package main

import (
    "fmt"
)

type Student struct{
    name string
    id string
}

func (st Student) printStudent(){
    st.id = "001"
    fmt.Println(st)
}

func (st *Student) pprintStudent(){
    st.id = "001"
    fmt.Println(st)
}

func main() {
    st1 := Student{"Chang", "0000"}
    fmt.Println(st1)
    st1.printStudent()
    fmt.Println(st1)
    st1.pprintStudent()
    fmt.Println(st1)
}

