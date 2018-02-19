package main

import (
	"fmt"
)

func main() {
	// Maps. Son los diccionarios o Hashtables ["clave", "valor"]

	// Declaraciones
	x := make(map[string]string)
	fmt.Println(x)

	y := make(map[string]string, 2)
	fmt.Println(y)

	// Agregar valores
	x["nombre"] = "Xavi"
	x["edad"] = "38"
	fmt.Println(x)

	// Acceder a los valores
	fmt.Println(x["nombre"])
	fmt.Println(x["edad"])

	// Declaracion 3
	edades := map[string]int{
		"Xavi":  38,
		"Elena": 35,
		"Lolo":  45,
		"Maria": 19,
	}

	fmt.Println(edades)
	fmt.Println(edades["Xavi"])

	// Para las claves se puede usar cualquier tipo compatible con ==
	// Para los valores sirve cuanquier tipo de dato
	dias := map[int]string{
		1: "Lunes",
		2: "Martes",
		3: "Miercoles",
		4: "Jueves",
		5: "Viernes",
		6: "Sabado",
		7: "Domingo",
	}

	fmt.Println(dias[2])

	// Eliminar elementos de un map
	delete(edades, "Maria")
	fmt.Println(edades)

	// Si el elemento no existe los amps devuelven 0 o el valor de inicializacion del tipo de datos p ej. nil
	fmt.Println("La edad de Xavi es: ", edades["xavi"])
	fmt.Println("Dia seleccionado es: ", dias[8])

	// Saber si un elemento existe en un map. El ok es un segundo valor que devuelve el map que es un booleano que indica si ha encontrado algo
	if edad, ok := edades["xavi"]; ok {
		if edad < 18 {
			fmt.Println("Menor de edad")
		} else {
			fmt.Println("Mayor de edad")
		}
	} else {
		fmt.Println("El valor no existe")
	}

	// Tamaño del map
	fmt.Println(len(edades))

	// Recorrer un map - Los maps son DESORDENADOS
	for index, dia := range dias {
		fmt.Printf("El num de dia %d es %q\n", index, dia)
	}
}
