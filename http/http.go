package main
import "net/http"
import "fmt"

func main() {
    res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
    fmt.Println(res)
}
