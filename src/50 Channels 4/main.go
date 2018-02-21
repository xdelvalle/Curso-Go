package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	bufferedChannel := make(chan string, 25)

	for i := 1; i < 5; i++ {
		go enviarMensaje(bufferedChannel, i)
	}

	imprimir(bufferedChannel)
}

func enviarMensaje(out chan<- string, numero int) {
	for {
		out <- "Mensaje: " + strconv.Itoa(numero)
		fmt.Println("Enviado mensaje func: ", numero)
	}
}

func imprimir(in <-chan string) {
	for x := range in {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}
