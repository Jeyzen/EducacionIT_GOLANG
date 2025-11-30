package main

import (
	"fmt"
)

func contar(canal chan int){
	x:= 0
	for {
		canal<-x
		x++
	}
}

func imprimir (canal chan int){
	var valor int
	for {
		valor = <-canal
		fmt.Println(valor)
	}
}

func main() {
    canal := make(chan int,3)
	go contar(canal)
	go imprimir(canal)
	//close(canal)
	var s string
	fmt.Scanln(&s)
}
