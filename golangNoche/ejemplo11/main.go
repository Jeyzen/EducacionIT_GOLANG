package main

import "fmt"




type Persona struct{
	Nombre string
	Edad int
}

func (p Persona) Presentarse(){
	fmt.Println("Hola soy",p.Nombre,"y tengo",p.Edad)
}

func main(){

	x := Persona{"Juan",32}
	x.Presentarse()
}