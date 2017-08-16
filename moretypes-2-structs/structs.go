package main

import "fmt"

type Cube struct {
    length int
}

func (c Cube) volume() int {
    return c.length * c.length * c.length
}

func (c Cube) surfaceArea() int {
    return c.length * c.length * 6
}

func main() {
    c1 := Cube{3}
    fmt.Println("Cube", c1)
    fmt.Println("Cube volume is: ", c1.volume())
    fmt.Println("Cube surface area is: ", c1.surfaceArea())
}
