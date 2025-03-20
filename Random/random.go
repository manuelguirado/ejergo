//generar slice aleatorio

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	mySlice := make([]int, 100)

	for i := 0; i < len(mySlice); i++ {
		mySlice[i] = rand.Intn(100)
	}
	fmt.Println((mySlice))
}
