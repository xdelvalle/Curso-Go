package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	done := time.After(20 * time.Second) // Es un channel, cuando pases 20 seg. se mandara el mensaje por "done"
	eco := make(chan []byte)

	go leerEntrada(eco)

	for {
		// Parecido al switch, pero evalua cuando a traves de un channel vienen datos
		select {
		case datos := <-eco:
			os.Stdout.Write(datos)
		case <-done:
			fmt.Println("Se termino el tiempo")
			os.Exit(0)
		}
	}
}

func leerEntrada(out chan<- []byte) {
	for {
		datos := make([]byte, 1024)
		leidos, _ := os.Stdin.Read(datos)

		if leidos > 0 {
			out <- datos
		}
	}
}
