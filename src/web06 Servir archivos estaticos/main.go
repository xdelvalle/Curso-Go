package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/", fs)

	// Con nuestro propio server, podemos modificar algunos datos
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening..")
	log.Fatal(server.ListenAndServe()) // Si da error se para el programa
}
