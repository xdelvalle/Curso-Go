package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Print("Esperando conexion al puerto #8000")
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		// manejarConexion(conn) // Manejamos una conexion a la vez.
		go manejarConexion(conn) // Manejamos varias conexiones
	}
}

func manejarConexion(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			log.Print("Cliente desconectado!")
			return // p ej. si el cliente de desconecta
		}

		time.Sleep(1 * time.Second)
	}
}
