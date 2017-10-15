package main
import "fmt"
func main() {
fmt.Print("Enter the size of array: ")
var n int
fmt.Scanf("%d", &n)
x := make([]float64, n)
for i := 0; i < len(x); i++ {
fmt.Printf("Enter the number at position %d : ", i+1)
fmt.Scanf("%f", &x[i])
}
var total float64 = 0
for _, value := range x {
total += value
}
fmt.Println( total / float64(len(x)))
}
