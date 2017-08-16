package main

import "fmt"

func main() {
    defer fmt.Println("world")
    defer fmt.Println("world2")
    return
    defer fmt.Println("world3")
    fmt.Println("hello")
}

