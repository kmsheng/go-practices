package main

import "fmt"

func main() {
    names := [4]string{
        "David",
        "Eric",
        "Alice",
        "Brady",
    }
    fmt.Println(names)

    a := names[0:2]    // David, Eric
    b := names[1:3]    // Eric, Alice, Brady

    b[0] = "deleted"
    fmt.Println(a, b)
    fmt.Println(names)
}
