package main
import "fmt"
func main() {
const x int = 5
var arr [x]float64
for i := 0; i < x; i++ {
fmt.Printf("Enter the number at position %d : ", i+1)
fmt.Scanf("%f", &arr[i])
}
var total float64 = 0
for _, value := range arr {
total += value
}
fmt.Println(total / float64(len(arr)))
}
