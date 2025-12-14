package models

import (
	"encoding/json"
	"net/http"
)

type DolarResponse struct {
	Venta float64 `json:"venta"`
}

func ObtenerCotizacion() (float64, error) {
	resp, err := http.Get("https://dolarapi.com/v1/dolares/oficial")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data DolarResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	return data.Venta, nil
}
