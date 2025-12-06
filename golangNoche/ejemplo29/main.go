package main

//https://pkg.go.dev/net/http

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hola desde Go %s", r.Method)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servicio andando en localhost:8080")
	http.ListenAndServe(":8080", nil)
}