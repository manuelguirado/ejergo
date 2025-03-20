package main
import (
	"fmt"

)
func celsiusToFahrenheitToCelsius( celsius float64) float64 {
	return (celsius * 9/5) + 32
}
func FahrenheitToCelsius(f float64) float64 {
	return (f - 32 )  * 5/9
}
func main() {
	
	
	fmt.Println("A que unidad quieres convertir (c) celsius (f) fahrenheit")
	var unidad string
	var celsius float64	
	var fahrenheit float64
	fmt.Scanf("%s" , &unidad)
	switch unidad {
	case "c" :
		
		fmt.Print("dime la tempertura en  fahrenheit: ")
		fmt.Scanf("%f" , &fahrenheit)
		fmt.Println("la temperatura de farhenheit a celsius es: ", FahrenheitToCelsius(fahrenheit))
	case "f" :
		fmt.Print("dime la temperatura en celsius: ")
		fmt.Scanf("%f" , &celsius)
		fmt.Println("la temperatura de celsius a fahrenheit es" , celsiusToFahrenheitToCelsius(celsius))

	default :
		fmt.Println("unidad no valida")
	}
}