package main

import (
	"fmt"
)

func main() {
	// Range. Es un foreach

	nombres := []string{
		"Alejandro",
		"Maria",
		"Pedro",
		"Carlos",
	}

	// Range devuelve indice y valor
	for index, nombre := range nombres {
		fmt.Printf("El nombre %q esta en el index %d\n", nombre, index)
	}

	fmt.Println("**************************************")

	// Deschamos el indice, no nos interesa
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}

	fmt.Println("**************************************")

	// quitando el '_' nos devuelve solo el indice
	for index := range nombres {
		fmt.Println(index)
	}

	fmt.Println("**************************************")

	cadena := "Hola Xavi, aqui haciendo un curso de Go"

	for index, letra := range cadena {
		fmt.Printf("La letra %q esta en el indice %d\n", letra, index)
	}
}
