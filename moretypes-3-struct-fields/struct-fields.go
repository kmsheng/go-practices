package main

import "fmt"

type Car struct {
    brand string
    color string
}

func main() {
    c := Car{"volkswagen", "white"}
    c.color = "black"

    fmt.Println(c)
}
