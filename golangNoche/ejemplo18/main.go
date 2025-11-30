package main

import (
	"fmt"
	"math/rand"
)


func contiene(arr []int, num int) bool{
	for _,v := range arr{
		if v == num{
			return true
		}
	}
	return false
}

func main(){

	var numeros []int

	for i:=1;i<=6;i++{
		n:= rand.Intn(46)
		//esta n en el array
		if !contiene(numeros,n){
			numeros = append(numeros,n)
		}else{
			i--
		}
		//guardo en el array
		fmt.Println(i,n)
		//if len(numeros) == 6{
		//	break
		//}
	}
	fmt.Println(numeros)
}