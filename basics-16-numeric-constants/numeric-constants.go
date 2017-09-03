package main

import (
    "fmt"
    "math"
)

const (
    Big = 1 << 100
)

func needInt(x int) int {
    return x + 1
}

// An untyped constant takes the type needed by its context.
func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Printf("The maximun value for float64 is %v\n", math.MaxFloat64)
    fmt.Printf("Type of needFloat(Big): %T\n", needFloat(Big))

    // overflows int
    // fmt.Printf("Type of needInt(Big): %T\n", needInt(Big))
}
