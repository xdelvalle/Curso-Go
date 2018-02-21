package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Router Gorilla Mux. Permite hacer mas cosas que el que viene de serie con Go
// Para instalar gorilla mux --> $go get github.com/gorilla/mux

func main() {
	r := mux.NewRouter().StrictSlash(false) //diferencia entre /api/user con /api/user/

	r.HandleFunc("/api/users", GETUsers).Methods("GET")
	r.HandleFunc("/api/users", POSTUsers).Methods("POST")
	r.HandleFunc("/api/users", PUTUsers).Methods("PUT")
	r.HandleFunc("/api/users", DELETEUsers).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening #8080..")
	log.Fatal(server.ListenAndServe())
}

func GETUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde el metodo GET")
}

func POSTUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde el metodo POST")
}

func DELETEUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde el metodo DELETE")
}

func PUTUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde el metodo PUT")
}
