package main

import (
	"fmt"
	"log"

	"tp_final/db"
	"tp_final/services"
)

func main() {
	db.InitDB()

	fmt.Println("Bienvenido a Hamburguesas IT")
	fmt.Print("Ingrese su nombre encargad@: ")

	var encargado string
	fmt.Scanln(&encargado)

	// Registrar IN
	err := services.RegistrarEntrada(encargado)
	if err != nil {
		log.Println("Error registrando entrada:", err)
	}

	for {
		fmt.Println("\n--- MENÚ PRINCIPAL ---")
		fmt.Println("1) Ingreso Nuevo Pedido")
		fmt.Println("2) Cambio de Turno")
		fmt.Println("3) Apagar Sistema")
		fmt.Print("Seleccione opción: ")

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			services.ProcesarPedido(encargado)

		case 2:
			encargado = services.CambiarTurno(encargado)

		case 3:
			services.RegistrarSalida(encargado)
			fmt.Println("Sistema apagado. ¡Hasta luego!")
			return

		default:
			fmt.Println("Opción inválida, intente nuevamente.")
		}
	}
}
