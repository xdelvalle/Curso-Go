package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup le indica al programa que debe esperar a que finalicen las gorutinas
var waitGroup sync.WaitGroup

func main() {
	// Espera a que finalicen 2 gorutinas
	waitGroup.Add(2)

	fmt.Println("Iniciamos las gorutinas..")

	go imprimirCantidad("A")
	go imprimirCantidad("B")

	fmt.Println("Esperando a que finalicen las gorutinas..")

	// Va preguntandose periodicamente si ya han finalizado todas las gorutinas, si es asi finaliza
	waitGroup.Wait()
	fmt.Println("FInalizado!")
}

func imprimirCantidad(etiqueta string) {
	// Le indicamos que la gorutina termina (el defer lo ejecuta al final de la funcion)
	defer waitGroup.Done()

	for i := 0; i < 10; i++ {
		// Espera aleatoria
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Cantidad: %d de %s\n", i, etiqueta)
	}
}
