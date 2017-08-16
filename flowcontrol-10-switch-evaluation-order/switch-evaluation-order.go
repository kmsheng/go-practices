package main

import "fmt"

func main() {
    // does not print anything becase case 0 is already matched
    i := 0
    switch i {
    case 0:
    case 1:
        fmt.Println("1")
    default:
        fmt.Println("default")
    }
}
