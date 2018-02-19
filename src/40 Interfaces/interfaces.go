package main

import (
	"fmt"
)

// Persona struct
type Persona struct {
	nombre string
	email  string
	edad   int
}

// Moderador struct
type Moderador struct {
	Persona
	Foro string
}

// Administrador struct
type Administrador struct {
	Persona
	Seccion string
}

// Nombre ..
func (p Persona) Nombre() string {
	return p.nombre
}

// Email ..
func (p Persona) Email() string {
	return p.email
}

// Edad ..
func (p Persona) Edad() int {
	return p.edad
}

// Presentarse ..
// func Presentarse(p Persona) {
// 	fmt.Println("Nombre: ", p.Nombre())
// 	fmt.Println("Email : ", p.Email())
// }

// Presentarse ..
func Presentarse(p Usuario) {
	fmt.Println("Nombre: ", p.Nombre())
	fmt.Println("Email : ", p.Email())
}

// PresentarseAdmin ..
func PresentarseAdmin(a Administrador) {
	fmt.Println("Nombre: ", a.Nombre())
	fmt.Println("Email : ", a.Email())
}

// PresentarseModerador ..
func PresentarseModerador(m Moderador) {
	fmt.Println("Nombre: ", m.Nombre())
	fmt.Println("Email : ", m.Email())
}

// AbrirForo ..
func (m Moderador) AbrirForo() {
	fmt.Println("Abrir un foro")
}

// CerrarForo ..
func (m Moderador) CerrarForo() {
	fmt.Println("Cerrar un foro")
}

// CrearForo ..
func (a Administrador) CrearForo() {
	fmt.Println("Crear un foro")
}

// EliminarForo ..
func (a Administrador) EliminarForo() {
	fmt.Println("Elimina un foro")
}

// Usuario Interfaz Usuario: No de codifican los metodos
type Usuario interface {
	Nombre() string
	Email() string
}

func main() {
	xavi := Persona{"Xavi", "xavi.delvalle@gmail.com", 38}
	Presentarse(xavi)

	juan := Moderador{Persona{"Juan", "j@j.com", 30}, "Juegos"}
	pedro := Administrador{Persona{"Pedro", "p@p.com", 50}, "PC"}

	// PresentarseModerador(juan)
	// PresentarseAdmin(pedro)

	Presentarse(juan)
	Presentarse(pedro)

	var i Usuario
	i = xavi
	fmt.Println("i: ", i)
	fmt.Println("i.Nombre(): ", i.Nombre())
	fmt.Println("i.Email(): ", i.Email())

	i = juan
	fmt.Println("i: ", i)
	fmt.Println("i.Nombre(): ", i.Nombre())
	fmt.Println("i.Email(): ", i.Email())
	//i.CerrarForo() // No se puede porque con la variable i SÓLO tiene acceso a los metodos declarados en la interfaz (aunque juan tiene mas)

}
