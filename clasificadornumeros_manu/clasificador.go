package main
import "fmt"
func espar( numero int) int {
	if numero% 2  == 0 {
		fmt.Println("El número es par")
	}else{
		fmt.Println("El número es impar")
	}
	return numero

}
func esPositivo( numero  int) int {
	if numero > 0 {
		fmt.Print("El número es positivo")
	}else if numero < 0 {
		fmt.Println("el numero es negativo")
	}else{
		fmt.Println("El número es 0")
	}
	return numero
}
func esMultiplo5 (numero int) int {
	if numero % 5 == 0 {
		fmt.Println("el numeroe es multiplo de 5")
	}else{
		fmt.Print("El número no es multiplo de 5")
	}
	return numero
}
func main() {
	var numero int
	fmt.Println("introduce un numero ")
	fmt.Scanf("%d" , &numero)
	fmt.Println("ingrese operacion a realizar (1) par (2 )  positivo (3) multiplo de 5")
	var operacion int
	fmt.Scanf("%d" , &operacion)
	switch operacion {
	case 1 :
		espar(numero)
	case 2: 
	    esPositivo(numero)
	case 3:
		esMultiplo5(numero)
	}
}