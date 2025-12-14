package db

import (
	"time"
)

func RegistrarEvento(nombre, evento string, caja float64) error {
	query := `INSERT INTO registro (encargado, fecha, evento, caja) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(query, nombre, time.Now().Format(time.RFC3339), evento, caja)
	return err
}
