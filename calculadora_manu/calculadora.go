package main

import "fmt"
func restar(a int, b int) int {
	return a -b 
}
func sumar(a int, b int) int {
	return a + b
}
func multiplicar(a int , b int)  int {
	return a  * b

}
func dividir(a int, b int) int {
	return a /  b
}
func resto(a int , b int)  int {
	return a % b
}
func main() {
	fmt.Println("ingrese un operador para hacer una operacion (+,-,*,/,%)")
	var operador string
	fmt.Scanf("%s" , &operador)
	fmt.Print("ingrese el primer numero: ")
	var numero1 int
	fmt.Scanf("%d" , &numero1)
	fmt.Print("ingrese el segundo numero")
	var numero2 int
	var resultado int
	fmt.Scanf("%d" ,&numero2)
	switch operador {
	case "+" :
		resultado = sumar(numero1,numero2)
		fmt.Println("el resultado es:" , resultado)
	case "-" :
		resultado = restar(numero1, numero1)
		fmt.Println("el resultado es:" , resultado)
	case "*":
		resultado = multiplicar(numero1, numero2)
		fmt.Println("el resultado es:" , resultado)
	case "/" :
		resultado = dividir(numero1,numero2)
		fmt.Println("el resultado es:" , resultado)
	case "%":
		resultado = resto(numero1,numero2)
		fmt.Println("el resultado es:" , resultado)
	default:
		fmt.Print("operador no valido")
		
	}
}