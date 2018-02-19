package main

import (
	"fmt"
	"os"
)

// Defer. Permite que la funcion se ejecute al terminar la ejecución del ámbito donde está declarada esa funcion.

func main() {
	// Se ejecurará cuando termine todo el main()
	defer primera()
	segunda()

	// Abrimos fichero
	f, err := os.Open("texto.txt")

	if err != nil {
		panic(err)
	}

	// AQUI CERRAMOS EL FICHERO CON DEFER (despues de que se ejecute toda la función)
	defer f.Close()

	// Slice para almacenar el fichero leido
	data := make([]byte, 175)

	// Leemos fichero
	bytesLeidos, err := f.Read(data)

	if err != nil {
		panic(err)
	}

	// Imprimimos lo leido
	fmt.Printf("Bytes leidos: %d\nTexto leido: %q\nerror: %v\n", bytesLeidos, data, err)

	// Si cerramos aquí es posible que si hay error nunca se cierre, por eso lo cerraremos con defer.
	//f.Close()

	// Si hay mas de un defer en la función se ejecutaran comom una cola LIFO el ultimo en ponerse sera el primero en salir.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func primera() {
	fmt.Println("Primera")
}

func segunda() {
	fmt.Println("Segunda")
}
