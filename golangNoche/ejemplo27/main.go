package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Estructura según el JSON de DolarAPI
type Dolar struct {
	Moneda            string    `json:"moneda"`
	Casa              string    `json:"casa"`
	Nombre            string    `json:"nombre"`
	Compra            float64   `json:"compra"`
	Venta             float64   `json:"venta"`
	FechaActualizacion time.Time `json:"fechaActualizacion"`
}

func main() {
	url := "https://dolarapi.com/v1/dolares"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error en GET: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status inesperado: %v", resp.Status)
	}

	// Decodificar JSON
	var valores []Dolar
	if err := json.NewDecoder(resp.Body).Decode(&valores); err != nil {
		log.Fatalf("error al decodificar JSON: %v", err)
	}

	// Mostrar resultados
	for _, d := range valores {
		fmt.Printf("=== %s (%s) ===\n", d.Nombre, d.Casa)
		fmt.Printf("Compra: %.2f | Venta: %.2f\n", d.Compra, d.Venta)
		fmt.Printf("Actualización: %s\n\n", d.FechaActualizacion.Format(time.RFC3339))
	}
}
