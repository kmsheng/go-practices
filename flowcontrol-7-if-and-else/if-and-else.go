package main

import "fmt"

func main() {
    const somethingWrong = true
    if somethingWrong {
        fmt.Println("Call your daddy here")
    } else {
        fmt.Println("Nothing happened")
    }
}
