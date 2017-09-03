package main

import "fmt"

type Ball struct {
	Radius uint
}

var m map[string]Ball

func main() {
	m = make(map[string]Ball)
	// save a ball with radis 5 to map
	m["My ball"] = Ball{5}
	fmt.Println(m["My ball"])
}
