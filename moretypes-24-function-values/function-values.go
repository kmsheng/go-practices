package main

import "fmt"

func compute(fn func(float64, float64) float64) float64 {
	return fn(2, 4)
}

func main() {
	myFunc := func(x, y float64) float64 {
		return x * y
	}
	fmt.Println(myFunc(2, 2))
	fmt.Println(compute(myFunc))
}
