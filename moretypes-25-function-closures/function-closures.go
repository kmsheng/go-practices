package main

import "fmt"

func createAddFunc(a, b int) func(x, y int) int {
	return func(x, y int) int {
		return a + b + x + y
	}
}

func main() {
	addFunc := createAddFunc(1, 2)
	fmt.Println(addFunc(3, 4))
}
