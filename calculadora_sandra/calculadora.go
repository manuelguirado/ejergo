package main

import "fmt"

// Funciones de suma, resta, multiplicacion y division
func suma(num1, num2 int) int {
	return num1 + num2
}

func resta(num1, num2 int) int {
	return num1 - num2
}

func division(num1, num2 int) int {
	return num1 / num2
}

func multiplicacion(num1, num2 int) int {
	return num1 * num2
}

func main() {
	var num1 int
	var num2 int

	// Solicitar entrada del usuario
	fmt.Println("Introduce el primer número:")
	fmt.Scanln(&num1)
	fmt.Println("Introduce el segundo número:")
	fmt.Scanln(&num2)

	// Llamar a las funciones de las distintas operaciones y almacenar el resultado
	resultado := suma(num1, num2)
	resultado1 := resta(num1, num2)
	resultado2 := multiplicacion(num1, num2)
	resultado3 := division(num1, num2)

	// Mostrar el resultado
	fmt.Println("El resultado de la suma es:", resultado)
	fmt.Println("El resultado de la resta es:", resultado1)
	fmt.Println("El resultado de la division es:", resultado2)
	fmt.Println("El resultado de la division es:", resultado3)
}
