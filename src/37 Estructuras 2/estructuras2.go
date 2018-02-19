package main

import (
	"fmt"
)

// Persona struct
type Persona struct {
	Nombre   string
	Apellido string
}

// Estudiante struct
type Estudiante struct {
	Persona
	Carrera string
}

// Profesor struct
type Profesor struct {
	Estudiante
	Carrera string
}

func main() {
	xavi := Estudiante{
		Persona{
			Nombre:   "Xavi",
			Apellido: "del Valle",
		},

		"Telecos",
	}

	fmt.Println("Xavi: ", xavi)

	// Accediendo a los campos, desde Estudiante hay acceso a todos los campos de Persona sin necesidad de pasar por xavi.Persona.Nombre

	fmt.Println("Nombre: ", xavi.Nombre)
	fmt.Println("Apellido: ", xavi.Apellido)
	fmt.Println("Carrera: ", xavi.Carrera)

	fmt.Println("********************************************")

	pepe := Profesor{
		Estudiante{
			Persona{
				Nombre:   "Pepe",
				Apellido: "Gotera",
			},
			"Informatica",
		},
		"Telecomunicaciones",
	}

	fmt.Println("Nombre: ", pepe.Nombre)
	fmt.Println("Apellido: ", pepe.Apellido)
	fmt.Println("Carrera: ", pepe.Carrera)
	fmt.Println("Carrera Estudiante: ", pepe.Estudiante.Carrera)

	fmt.Println("********************************************")

	// Declarando una variable de tipo Profesor en una linea
	manuel := Profesor{Estudiante{Persona{"Manuel", "Diaz"}, "ADE"}, "Industriales"}
	fmt.Println("Manuel: ", manuel)

	fmt.Println("********************************************")

	// Otra forma
	var jose Profesor
	jose.Nombre = "Jose"
	jose.Apellido = "Mota"
	jose.Estudiante.Carrera = "MBA"
	jose.Carrera = "Econonomia"

	fmt.Println("Jose: ", jose)

}
