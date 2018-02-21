package main

import (
	"fmt"
	"net/http"
)

func main() {
	msg := mensaje{
		msg: "Hola Mundo desde el manejador",
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", holaMundoHandler)
	mux.HandleFunc("/prueba", pruebaHandler)
	mux.HandleFunc("/usuario", usuarioHandler)

	// Ejecutamos desde nuestro Handle msg
	mux.Handle("/hola", msg)

	http.ListenAndServe(":8080", mux)
}

type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
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
