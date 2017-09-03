package main

import "fmt"

type Paper struct {
	distance uint
}

func main() {
	// note: the type of each element is omitted
	var m = map[string]Paper{
		"paper1": {5},
		"paper2": {10},
	}
	fmt.Println(m)
}
