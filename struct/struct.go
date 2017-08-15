package main

import "fmt"

type person struct {
    name string
    age int
}

func main() {
    fmt.Println(person{"kmsheng", 20})
    fmt.Println(person{"alice", 30})

    anotherPerson := person{"john", 15}

    anotherPerson.age = 100

    fmt.Println(anotherPerson)
}
