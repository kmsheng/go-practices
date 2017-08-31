package main

import "fmt"

func split(sum int) (first, second int) {
    first = sum * 1
    second = sum * 2
    // "naked" return, should used only in short functions
    return
}

func main() {
    fmt.Println(split(100))
}
