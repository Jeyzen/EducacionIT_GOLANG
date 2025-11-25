package main

import "fmt"


/*

Resumen simple (para enseñar)

	Un puntero almacena una dirección de memoria.
	&x → dirección de x.
	*p → valor apuntado por p.
	Usá punteros para:
		Modificar variables en funciones.
		Evitar copias grandes.
		Modelar estructuras complejas.
*/

type Persona struct{
	Nombre string
	Edad int
}

func cambiarEdad(p* Persona){
	p.Edad = 39
}


func main(){
	w := new(int)
	*w = 77
	fmt.Println(*w,w) // 0
	x := Persona{"Carlos",34}
	cambiarEdad(&x)
	fmt.Println(x)

	//
}