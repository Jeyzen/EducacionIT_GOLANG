package main

import (
	"fmt"
	"strconv"
)

func main() {
	numeroUno := "48"
	numeroDos := "52"
	var x1, x2 int
	x1, _ = strconv.Atoi(numeroUno) // 48,nil
	x2, _ = strconv.Atoi(numeroDos) // 52,nil
	suma := x1 + x2

	fmt.Print("La suma es ", suma)
	fmt.Printf(" %T \n", suma)

	texto := strconv.Itoa(suma)
	fmt.Print("La suma es ", texto)
	fmt.Printf(" %T \n", texto)

	pit := "3.141592"
	pin, _ := strconv.ParseFloat(pit, 64) // FormatFloat

	fmt.Print("Pi", pin)
	fmt.Printf(" %T \n", pin)

}
