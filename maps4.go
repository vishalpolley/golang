package main
import "fmt"
func main() {
elements := map[string]string{
"H" : "Hydrogen",
"He": "Helium",
"Li": "Lithium",
"Be": "Beryllium",
"B" : "Boron",
"C" : "Carbon",
}
if el, ok := elements["Un"]; ok {
fmt.Println(el, ok)
}
}
