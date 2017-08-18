package main

import "fmt"

func main() {
    var scores = []int{100, 20, 30, 74, 86}
    for score := range scores {
        fmt.Printf("score: %d\n", score)
    }
}
