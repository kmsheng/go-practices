package main

import "fmt"

func main() {
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s)

    s = s[:0]
    printSlice(s)

    s = s[:4]
    printSlice(s)

    s = s[2:]
    printSlice(s)

    // panic: runtime error: slice bounds out of range
    // s = s[:100]
    // printSlice(s)

    // invalid slice index -100 (index must be non-negative)
    // s = s[-100:]
    // printSlice(s)
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
