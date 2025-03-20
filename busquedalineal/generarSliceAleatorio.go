package main
import (
	"math/rand"
	"fmt"
)
func generarSlice (tamano int) []int {
	var slice = make([]int,tamano)
	for i := 0; i < tamano; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}
func main() {
	var tamano int
	fmt.Print("ingrese un tamano para el slice")
	fmt.Scanf("%d " , &tamano)
	fmt.Println(generarSlice(tamano))

}