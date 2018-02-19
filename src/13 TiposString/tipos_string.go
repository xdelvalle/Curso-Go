package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Los String son una secuencia de Bytes (Slice de Bytes)
	// Son indexables
	// Son inmutables

	var cadena string
	cadena = "Hola Mundo"
	fmt.Println("Cadena: ", cadena)

	fmt.Println("Lenght: ", len(cadena))

	fmt.Println("cadena[2] = ", cadena[2]) // 108 corresponde a la "l"
	fmt.Println("cadena[0:4] = ", cadena[0:4])

	cadena = cadena + " nuevamente"
	fmt.Println(cadena)

	cadena += " siiiiiiiiiii"
	fmt.Println(cadena)

	// Otra forma de declarar strings es con los acentos abiertos. Toma la cadena literal con espacios incluidos, tabs y caracteres especiales
	cadena = `
	<html>
		<head>
			<meta charset="utf-8">
			<title></title>
		</head>
		<body>

		</body>
	</html>
	`
	fmt.Println(cadena)

	cadena = "Hola Mundo \"25\"" // Usar comillas dentro de un string con \ delante
	fmt.Println(cadena)

	edad := 38
	cadena = "la edad es " + strconv.Itoa(edad)
	fmt.Println(cadena)
}
