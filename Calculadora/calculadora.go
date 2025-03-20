// Calculadora simple
package main

import (
	"fmt"
)

func main() {
	var num1 int
	fmt.Println("Introduce un número")
	fmt.Scanf("%d", &num1)

	var num2 int
	fmt.Println("Introduce otro número")
	fmt.Scanf("%d", &num2)

	total := num1 + num2

	fmt.Println("El resultado es ", total)
}
