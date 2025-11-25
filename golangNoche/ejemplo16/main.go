package main

import (
	"fmt"
)

func mostrar0() {
	for f := 1; f <= 1000; f++ {
		fmt.Print("0-")
	}
}

func mostrar1() {
	for f := 1; f <= 1000; f++ {
		fmt.Print("1-")
	}
}

func main() {
	go mostrar0()
	go mostrar1()
	continua := ""
	fmt.Scan(&continua)
}
