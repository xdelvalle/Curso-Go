package main

// import (
// 	"os"
// 	"text/template"
// )

// const tp = `Nombre: {{.Nombre}} Edad: {{.Edad}}`

// // Persona ..
// type Persona struct {
// 	Nombre string
// 	Edad   int
// }

// func main() {
// 	persona := Persona{"Xavi", 38}

// 	t := template.New("persona")

// 	t, err := t.Parse(tp)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = t.Execute(os.Stdout, persona)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// ****** RANGE **************************************
// import (
// 	"os"
// 	"text/template"
// )

// // Template que espera varias Personas
// const tp = `{{range .}}
// Nombre: {{.Nombre}} Edad: {{.Edad}}
// {{end}}`

// // Persona ..
// type Persona struct {
// 	Nombre string
// 	Edad   int
// }

// func main() {
// 	persona := []Persona{
// 		{"Xavi", 38},
// 		{"Pepe", 18},
// 		{"Laia", 3},
// 		{"Marc", 68},
// 		{"Mia", 33},
// 	}

// 	t := template.New("persona")

// 	t, err := t.Parse(tp)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = t.Execute(os.Stdout, persona)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// *********** IF *****************************************
import (
	"os"
	"text/template"
)

// El if sirve para el valor 0 de cada tipo .. int - 0, string - nil etc..
const tp = `
{{range .}}
	{{if .Edad}} 
		Nombre: {{.Nombre}} - Edad: {{.Edad}} - Correcto
	{{else}}
		Nombre: {{.Nombre}} - Edad: {{.Edad}} - INCORRECTO
	{{end}}
{{end}}`

// Persona ..
type Persona struct {
	Nombre string
	Edad   int
}

func main() {
	persona := []Persona{
		{"Xavi", 38},
		{"Pepe", 18},
		{"Laia", 3},
		{"Marc", 0},
		{"Mia", 33},
	}

	t := template.New("persona")

	t, err := t.Parse(tp)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, persona)
	if err != nil {
		panic(err)
	}
}
