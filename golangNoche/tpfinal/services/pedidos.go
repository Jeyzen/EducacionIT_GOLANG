package services

import (
	"fmt"
	"log"

	"tp_final/db"
	"tp_final/models"
)

func ProcesarPedido(encargado string) {
	var venta models.Venta

	fmt.Print("Nombre del cliente: ")
	fmt.Scanln(&venta.Cliente)

	fmt.Print("Cantidad Combo Simple: ")
	fmt.Scanln(&venta.ComboS)

	fmt.Print("Cantidad Combo Doble: ")
	fmt.Scanln(&venta.ComboD)

	fmt.Print("Cantidad Combo Triple: ")
	fmt.Scanln(&venta.ComboT)

	fmt.Print("Cantidad Flurby: ")
	fmt.Scanln(&venta.Flurby)

	venta.TotalUSD = float64(venta.ComboS)*5 + float64(venta.ComboD)*6 +
		float64(venta.ComboT)*7 + float64(venta.Flurby)*2

	cot, err := models.ObtenerCotizacion()
	if err != nil {
		log.Println("Error consultando API dólar:", err)
		return
	}
	venta.Cotizacion = cot
	venta.TotalARS = venta.TotalUSD * cot

	fmt.Printf("Total en USD: %.2f\n", venta.TotalUSD)
	fmt.Printf("Cotización dólar: %.2f\n", cot)
	fmt.Printf("Total en ARS: %.2f\n", venta.TotalARS)

	fmt.Print("Ingrese monto abonado en ARS: ")
	var pago float64
	fmt.Scanln(&pago)

	vuelto := pago - venta.TotalARS
	fmt.Printf("Vuelto: %.2f ARS\n", vuelto)

	fmt.Print("¿Confirma pedido (Y/N)?: ")
	var conf string
	fmt.Scanln(&conf)

	if conf == "Y" || conf == "y" {
		err := db.InsertVenta(venta)
		if err != nil {
			log.Println("Error guardando venta:", err)
			return
		}
		fmt.Println("Venta registrada con éxito.")
	} else {
		fmt.Println("Venta cancelada.")
	}
}
