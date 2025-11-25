package main

import (
	"fmt"
	"strings" // ToUpper - ToLower - Replace - Repeat 
)

/*
Ingresar dos palabras por teclado. Verificar si la segunda
palabra esta contenida en la primera e indicar en que posición comienza
*/

func main() {
	var palabra1, palabra2 string
	fmt.Println("Ingrese primer palabra: ")
	fmt.Scan(&palabra1)
	fmt.Println("Ingrese segunda palabra: ")
	fmt.Scan(&palabra2)
	pos := strings.Index(palabra1, palabra2)
	if pos == -1 {
		fmt.Println("No esta contenida la 2 en la 1")
	} else {
		fmt.Println("La palabra", palabra2, "esta contenida en", palabra1)
	}

}
