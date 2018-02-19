package main

import (
	"fmt"
)

func main() {
	// Estructura de control switch

	var dia int

	fmt.Println("Inserta un dia de la semana")
	fmt.Scanln(&dia)

	// if dia == 1 {
	// 	fmt.Println("Lunes")
	// } else if dia == 2 {
	// 	fmt.Println("Martes")
	// } else if dia == 3 {
	// 	fmt.Println("Miercoles")
	// } else if dia == 4 {
	// 	fmt.Println("Jueves")
	// } else if dia == 5 {
	// 	fmt.Println("Viernes")
	// } else if dia == 6 {
	// 	fmt.Println("Sabado")
	// } else if dia == 7 {
	// 	fmt.Println("Domingo")
	// } else {
	// 	fmt.Println("Error, [1..7]")
	// }

	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Error, [1..7]")
	}

	numero := 26
	switch {
	case numero > 23:
		fmt.Println("El numero es mayor que 23")
		fallthrough // Es cuando queramos que siga evaluando los siguientes cases, si a la primera sale del switch (no existen breaks como en java)
	case numero > 25:
		fmt.Println("El numero es mayor que 25")
	default:
		fmt.Println("Almenos es un numero..")
	}
}
