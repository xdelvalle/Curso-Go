package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Pasamos las funciones como parametro, sin parentesis
	mux.HandleFunc("/", holaMundoHandler)
	mux.HandleFunc("/prueba", pruebaHandler)
	mux.HandleFunc("/usuario", usuarioHandler)

	http.ListenAndServe(":8080", mux)
}

func holaMundoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola mundo Handler!</h1>")
}

func pruebaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola mundo Handler! #Prueba</h1>")
}

func usuarioHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola mundo Handler! Usuario</h1>")
}
