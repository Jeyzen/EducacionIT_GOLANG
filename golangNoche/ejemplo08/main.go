package main

import "fmt"


func sumar10(a *int){
	*a = *a + 10
}


func main(){
	x := 10
	p := &x // variable p esa p apunta a un lugar de Memo. Cual? al lugar de memoria de x
	
	fmt.Println("p",p,"&x",&x)
	fmt.Println("x",x,"*p",*p)

	valor := 5
	sumar10(&valor)

	fmt.Println(valor)
}