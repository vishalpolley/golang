package main
import "fmt"
func main() {
x := make(map[string]int) // x is a map of strings to ints
x["key"] = 10
fmt.Println(x["key"])
}
