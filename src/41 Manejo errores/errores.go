package main

import (
	"errors"
	"fmt"
)

func main() {
	// err := baneado("miguel")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	// err = baneado("carlos")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	// err = baneado("pedro")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	// err = baneado("juan")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	// err = baneado("xxx")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	// *************************************************

	err := baneado("miguel")
	checkError(err)

	err = baneado("carlos")
	checkError(err)

	err = baneado("pedro")
	checkError(err)

	err = baneado("juan")
	checkError(err)

	err = baneado("xxx")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

var (
	// ErrorUsuarioNoValido ...
	ErrorUsuarioNoValido = errors.New("El usuario no es valido")
	// ErrorUsuarioEnProceso ...
	ErrorUsuarioEnProceso = errors.New("Usuario en proceso de registro")
	// ErrorPorDefecto ...
	ErrorPorDefecto = errors.New("Algo malo paso xD")
)

func baneado(usuario string) (err error) {
	ban := false

	switch usuario {
	case "miguel":
		ban = true
	case "carlos":
		ban = false
	case "juan":
		//return fmt.Errorf("El usuario no es valido")
		return ErrorUsuarioNoValido
	case "pedro":
		//return fmt.Errorf("Usuario en proceso de registro")
		return ErrorUsuarioEnProceso
	default:
		//return fmt.Errorf("Algo malo paso xD")
		return ErrorPorDefecto
	}

	if !ban {
		fmt.Printf("Usuario %s no esta baneado\n", usuario)
	} else {
		fmt.Printf("Usuario %s esta baneado\n", usuario)
	}

	return nil
}
