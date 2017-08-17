package main

import "fmt"

func main() {

    var arr[5]string
    arr[0] = "Hello"
    arr[1] = "World"

    anotherArr := [3]int{2, 3, 4}

    fmt.Println(arr, anotherArr)
}
