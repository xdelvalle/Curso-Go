package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// // Enteros con signo
	// var entero8 int8   // 8 bits con signo (-128 a 127)
	// var entero16 int16 // 16 bits con signo (-32768 a 32767)
	var entero32 int32 // 32 bits con signo (-2147483648 a 2147483647)
	//var entero64 int64 // 64 bits con signo (-9223372036854775808 a 9223372036854775807)

	// // Enteros sin signo
	// var uentero8 uint8   // 8 bits sin signo (0 a 255)
	// var uentero16 uint16 // 16 bits sin signo (0 a 65535)
	// var uentero32 uint32 // 32 bits sin signo (0 a 4294967295)
	// var uentero64 uint64 // 64 bits sin signo (0 a 18446744073709551615)

	// // Alias
	// var enteroByte byte // Alias para uint8
	// var enteroRune rune // Alias para int32

	// // Tipos dependientes de la plataforma
	// var enteroUint uint       // 32 o 64 bits
	var enteroInt int // 32 o 64 bits
	// var enteroUintptr uintptr // Entero sin signo grade para contener el valor de un puntero

	// Conversion de tipos

	// Suma int32 y int64

	entero32 = 25
	// entero64 = 85
	// fmt.Println(entero32 + int32(entero64))

	// Suma int32 y int
	enteroInt = 56
	fmt.Println(entero32 + int32(enteroInt))

	fmt.Println(unsafe.Sizeof(enteroInt), unsafe.Sizeof(entero32))
}
