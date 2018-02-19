package main

import (
	"fmt"
)

func main() {
	// Copy - Copia el contenido de un slice en otro slice

	origen := []int{1, 2, 3}
	destino := []int{3, 4, 5}

	copy(destino, origen)

	fmt.Println(origen, destino)

	origen2 := []int{1, 2, 3}
	destino2 := make([]int, 2)
	copy(destino2, origen2)
	fmt.Println(origen2, destino2)

	origen3 := []int{1, 2}
	destino3 := []int{3, 4, 5}
	copy(destino3, origen3)
	fmt.Println(origen3, destino3)
}
