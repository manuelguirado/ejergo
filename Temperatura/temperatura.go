// Conversor de temperatura

package main

import (
	"fmt"
)

func conversorCelsiusFarenheit(celsius float64) float64 {
	return ((celsius * 9.0) / 5.0) + 32
}
func conversorFarenheitCelsius(Farenheit float64) float64 {
	return ((Farenheit - 32.0) * 5.0) / 9.0
}

func main() {
	var temperatura float64
	fmt.Println("Introduce Temperatura Celsius para pasar a Farenheit")
	fmt.Scanf("%f", &temperatura)
	fmt.Println("Resultado: ", conversorCelsiusFarenheit(temperatura))

	fmt.Println("Introduce Temperatura Farenheit para pasar a Celsius")
	fmt.Scanf("%f", &temperatura)
	fmt.Println("Resultado: ", conversorFarenheitCelsius(temperatura))

}
