package main

import "fmt"
import "reflect"

type Student struct{
    name string
    id string
}
func main() {
    var a Student
    a.name, a.id = "Chang", "00000"
    fmt.Println(a)
    setID(a, "000010")
    fmt.Println(a)

    vsetID(&a, "000010")
    fmt.Println(a)


    b := new(Student)
    vsetID(b, "11111")
    fmt.Println(b)

    c := &Student{name: "Chang", id: "2222"}
    vsetID(c, "33333")
    fmt.Println(c)


    d := Student{name: "Chang", id: "2222"}
    setID(d, "4444")
    fmt.Println(d)
    vsetID(&d, "5555")
    fmt.Println(d)
    fmt.Println(reflect.TypeOf(d))
}

func setID(a Student, newID string){
    a.id = newID
    fmt.Println(a)
}

func vsetID(a *Student, newID string){
    (*a).id = newID
    fmt.Println(a)
}