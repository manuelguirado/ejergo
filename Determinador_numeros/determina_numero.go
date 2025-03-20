// Pide un número y determina si es:
// - Positivo/Negativo/Cero
// - Par/Impar
// - Múltiplo de 5

package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Introduce un número")
	fmt.Scanf("%d", &num)

	if num > 0 {
		fmt.Println("El número es positivo")
		if num%2 == 0 {
			fmt.Println("El número es par")
		} else {
			fmt.Println("El número es impar")
		}
		if num%5 == 0 {
			fmt.Println("El número es múltiplo de 5")
		}
	} else if num < 0 {
		fmt.Println("El número es negativo")
	} else {
		fmt.Println("El número es 0")
	}

}
