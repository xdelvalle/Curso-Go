package main

import (
	"fmt"
	"strings"
)

func main6() { // cambiar por main()
	cadena := "aeiou"
	cadena = strings.Map(func(r rune) rune {
		return r + 1 // Sumamos codigos UTF8
	}, cadena)

	fmt.Println(cadena)
}
