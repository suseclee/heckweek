package main

import (
    "fmt"
)

func main() {
    array := [3]int{}
    fmt.Println(array)
    println("############## ARRAY ##############")
    //Array
    array = [...]int{0,1,2}
    fmt.Println(array)

    changArray(array)
    fmt.Println(array)

    V_changArray(&array)
    fmt.Println(array)

    println("############## SLICE ##############")
    //Slice
    letters := []string{"a", "b", "c", "d"}
    fmt.Println(letters)

    changSlice(letters)
    fmt.Println(letters)

    V_changSlice(&letters)
    fmt.Println(letters)

    s := make([]byte, 5)
    fmt.Println(s)

    slice := array[:]
    fmt.Println(slice)

}

func changArray(a [3]int){
    a[0]= 6
}

func V_changArray(a *[3]int){
    (*a)[0]= 6
}

func changSlice(a []string){
    a[0]= "6"
}

func V_changSlice(a *[]string){
    (*a)[0]= "10"
}