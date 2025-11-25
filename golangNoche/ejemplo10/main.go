package main

import "fmt"

/*
Ejercicio A
Crear una función que:
Reciba un puntero a un número entero. Lo multiplique por 3.
Mostrar el valor antes y después de llamar a la función.


Ejercicio B
Crear una función intercambiar(a, b *int) que:
Reciba dos punteros a enteros. Intercambie sus valores (swap).
Probarla en main mostrando los valores antes y después.

*/

// Solucion Ejercicio A
func multiplicar(puntero *int) {
	*puntero = *puntero * 3 // perfecto
}

// Solucion Ejercicio B
func intercambiar(a, b *int) {
	aux := *a
	*a = *b
	*b = aux
}

// Solucion Ejercicio B , solucion de un compañero
func intercambiarE(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	//x := 10
	//multiplicar(&x)
	//fmt.Println(x)
	x := 10
	y := 5
	fmt.Println("x=", x, "y=", y)
	intercambiar(&x, &y)
	fmt.Println("x=", x, "y=", y)
}
