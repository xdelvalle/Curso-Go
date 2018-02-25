package main

import (
	"html/template"
	"os"
)

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
// import (
// 	"os"
// 	"text/template"
// )

// // El if sirve para el valor 0 de cada tipo .. int - 0, string - nil etc..
// const tp = `
// {{range .}}
// 	{{if .Edad}}
// 		Nombre: {{.Nombre}} - Edad: {{.Edad}} - Correcto
// 	{{else}}
// 		Nombre: {{.Nombre}} - Edad: {{.Edad}} - INCORRECTO
// 	{{end}}
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
// 		{"Marc", 0},
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

// ****** NOMBRADOS ************************************
// import (
// 	"os"
// 	"text/template"
// )

// // Persona ..
// type Persona struct {
// 	Nombre string
// 	Edad   int
// 	Pais   string
// }

// func mainX() {
// 	persona := Persona{"Xavi", 38, "BCN"}
// 	t := template.New("persona")

// 	//t, err := t.ParseFiles("templates/residentes.txt","templates/visitantes.txt") //Los ficheros separados por comas
// 	t, err := t.ParseGlob("templates/*.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = t.ExecuteTemplate(os.Stdout, "visitantes", persona)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// ********* TEMPLATES CON WITH ************************
const HERO = `
Hero Name: {{.Name}}
{{range .Emails}}
Email: {{.}}
{{end}}
{{with .Friends}}
{{range .}}
Friend Name: {{.Name}}
{{end}}
{{end}}
`

type Friend struct {
	Name string
}

type Hero struct {
	Name    string
	Emails  []string
	Friends []Friend
}

func main() {
	f1 := Friend{"Thor"}
	f2 := Friend{"Hulk"}

	t := template.New("Hero")
	t, err := t.Parse(HERO)

	if err != nil {
		panic(err)
	}

	hero := Hero{
		Name:    "Ironman",
		Emails:  []string{"ironmain@gmail.com", "turealironman@gmail.com"},
		Friends: []Friend{f1, f2},
	}

	err = t.Execute(os.Stdout, hero)
	if err != nil {
		panic(err)
	}
}
