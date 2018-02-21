package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	msg := mensaje{
		msg: "Hola Mundo desde el manejador",
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", holaMundoHandler)
	mux.HandleFunc("/prueba", pruebaHandler)
	mux.HandleFunc("/usuario", usuarioHandler)
	mux.Handle("/hola", msg)

	// COn nuestro propio server, podemos modificar algunos datos
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB (se podria poner el numero largo)
	}
	log.Println("Listening..")
	server.ListenAndServe()
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
