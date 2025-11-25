package main	

import (
	"fmt"
	"time"
)

func main(){
	var ahora time.Time
	ahora = time.Now()

	fmt.Println("Año ", ahora.Year())
	fmt.Println("Mes ",ahora.Month())
	fmt.Println("Dia ", ahora.Day())
	//seguir con horas, minutos y segundos

	navidad2025 := time.Date(2025,12,25,0,0,0,0,time.UTC)
	var actual time.Time
	actual = time.Now()
	fmt.Println(actual)
	diferencia := navidad2025.Sub(actual)
	fmt.Println("Cantidad de horas que faltan ",diferencia)
	for i:=0;i<10;i++{
		fmt.Println(i)
		time.Sleep(1*time.Second)
	}
}