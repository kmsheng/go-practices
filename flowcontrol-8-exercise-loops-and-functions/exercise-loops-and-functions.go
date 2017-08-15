package main

import (
    "fmt"
	"math"
)

func Sqrt(x float64) float64 {

	v := 1.0
    for {
        nextV := 0.5 * (v + (x / v))
        if math.Abs(v - nextV) < 0.00000000001 {
            return v
        }
        v = nextV
    }
}

func main() {
    fmt.Println(Sqrt(2))
}
