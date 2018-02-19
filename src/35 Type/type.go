package main

import (
	"fmt"
)

// Type, tipos de datos personalizados

// Dinero Type declarado por nosotros
type Dinero int

// Creamos el metodo String() (con mayuscula) para el tipo Dinero -- Todos los tipos tienen su metodo String() -  En este caso cada vez que se imprima algo tipo Dinero ira un $ delante
func (d Dinero) String() string {
	return fmt.Sprintf("$%d", d)
}

func main() {
	var sueldo Dinero
	sueldo = 25000
	fmt.Println("Sueldo: ", sueldo)

	aumento := 10000
	// sueldo += aumento // Operacion invalida, no son los mismos tipos
	sueldo += Dinero(aumento) // Hacemos la conversion

	fmt.Println("Sueldo + Aumento: ", sueldo)
}
