package services

import (
	"fmt"
	"log"

	"tp_final/db"
)

func RegistrarEntrada(encargado string) error {
	return db.RegistrarEvento(encargado, "IN", 0)
}

func RegistrarSalida(encargado string) {
	total, err := db.TotalPorEncargado(encargado)
	if err != nil {
		log.Println("Error calculando total:", err)
		total = 0
	}

	db.RegistrarEvento(encargado, "OUT", total)
	fmt.Printf("Turno cerrado. Caja acumulada: %.2f ARS\n", total)
}

func CambiarTurno(encargadoActual string) string {
	RegistrarSalida(encargadoActual)

	fmt.Print("Nuevo encargad@: ")
	var nuevo string
	fmt.Scanln(&nuevo)

	RegistrarEntrada(nuevo)
	return nuevo
}
