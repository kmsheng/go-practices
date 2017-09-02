package main

import (
    "fmt"
)

var (
    boolVar bool = false
    strVar string = "Hello World"
    uintVar uint = 1 << 64 - 1
    uint8Var uint8 = 1 << 8 - 1
)

func main() {
    fmt.Printf("Type: %T Value: %v\n", boolVar, boolVar)
    fmt.Printf("Type: %T Value: %v\n", strVar, strVar)
    fmt.Printf("Type: %T Value: %v\n", uintVar, uintVar)
    fmt.Printf("Type: %T Value: %v\n", uint8Var, uint8Var)
}
