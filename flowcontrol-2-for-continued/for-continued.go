package main

import "fmt"

func main() {
    sum := 1
    // while loop in golang
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)
}
