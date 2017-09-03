package main

import (
    "fmt"
)

func main() {
    var num = float64(100)
    fmt.Printf("type before converted -> %T\n", num)
    fmt.Printf("type after converted -> %T\n", uint(num))
}
