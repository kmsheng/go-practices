package main

import "fmt"

func main() {
	type Paper struct {
		distance uint
	}
	var m = map[string]Paper{
		"paper1": Paper{5},
		"paper2": Paper{10},
	}
	fmt.Println(m)
}
