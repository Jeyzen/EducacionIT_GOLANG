package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/PuerkitoBio/goquery"
)

func main() {
    url := "https://iol.invertironline.com/titulo/cotizacion/BCBA/GGAL/GRUPO-FINANCIERO-GALICIA-S.A/"

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Error request:", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        log.Fatalf("Status error: %v", resp.Status)
    }

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        log.Fatal("Error al parsear HTML:", err)
    }

    // Buscar exactamente el último precio
    precio := doc.Find(`span[data-field="UltimoPrecio"]`).First().Text()

    if precio == "" {
        fmt.Println("❌ No se encontró el precio en el HTML")
        return
    }

    fmt.Println("📈 Último precio:", precio)
}
