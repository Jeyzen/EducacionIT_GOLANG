package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./hamburgueseria.db")
	if err != nil {
		log.Fatal(err)
	}

	createVentas := `
	CREATE TABLE IF NOT EXISTS ventas (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		cliente TEXT,
		fecha TEXT,
		comboS INTEGER,
		comboD INTEGER,
		comboT INTEGER,
		flurby INTEGER,
		totalUSD REAL,
		totalARS REAL,
		cotizacion REAL
	);`

	createRegistro := `
	CREATE TABLE IF NOT EXISTS registro (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		encargado TEXT,
		fecha TEXT,
		evento TEXT,
		caja REAL
	);`

	DB.Exec(createVentas)
	DB.Exec(createRegistro)
}
