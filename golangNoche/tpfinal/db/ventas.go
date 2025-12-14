package db

import (
	"time"

	"tp_final/models"
)

func InsertVenta(v models.Venta) error {
	query := `INSERT INTO ventas 
	(cliente, fecha, comboS, comboD, comboT, flurby, totalUSD, totalARS, cotizacion) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := DB.Exec(query, v.Cliente, time.Now().Format(time.RFC3339),
		v.ComboS, v.ComboD, v.ComboT, v.Flurby, v.TotalUSD, v.TotalARS, v.Cotizacion)

	return err
}

func TotalPorEncargado(nombre string) (float64, error) {
	query := `SELECT COALESCE(SUM(totalARS), 0) FROM ventas 
	WHERE cliente IN 
	(SELECT cliente FROM ventas WHERE fecha >= 
		(SELECT MAX(fecha) FROM registro WHERE encargado = ? AND evento = "IN")
	)`

	var total float64
	err := DB.QueryRow(query, nombre).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
