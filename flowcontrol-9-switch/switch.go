package main

import (
    "fmt"
)

func main() {
    switch name := "Jeffery"; name {
    case "Jeffery":
        fmt.Println("It's Jeffery")
        fallthrough
    case "David":
        fmt.Println("It's David")
        fallthrough
    default:
        fmt.Println("default")
    }
}

