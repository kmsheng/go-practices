package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["key"] = 10
	fmt.Println("value -> ", m["key"])

	m["key"] = 20
	fmt.Println("value -> ", m["key"])

	delete(m, "key")
	fmt.Println("value -> ", m["key"])

	v, present := m["key"]
	fmt.Println("value", v, "present?", present)


	// value is 0 for getting a non-existent key
	fmt.Println("value for a key that does not exist in map -> ", m["whatever key that does not exist"])
}
