package main

// SERVERMUX. Es un enrutador de peticiones http

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hola mundo ServeMux!</h1>")
	})

	mux.HandleFunc("/prueba", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hola mundo ServeMux! #Prueba</h1>")
	})

	mux.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hola mundo ServeMux! #Usuario</h1>")
	})

	http.ListenAndServe(":8080", mux)
}
