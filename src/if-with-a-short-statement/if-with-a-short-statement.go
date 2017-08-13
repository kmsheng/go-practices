package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    fmt.Println("V", math.Pow(x, n), "lim", lim)
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}

func main() {
    fmt.Println(
        pow(3, 3, 20),
    )
}

