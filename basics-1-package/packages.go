package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    fmt.Println("My favorite number is", r.Intn(19))
}
