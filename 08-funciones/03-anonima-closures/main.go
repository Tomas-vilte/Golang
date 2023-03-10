package main

import (
	"fmt"
	"strings"
)

// closures
func repeat(n int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {
	repeat3 := repeat(3)
	fmt.Println(repeat3("Hola"))
	fmt.Println(repeat3("Mundo"))
	// Funcion anonima
	/*func() {
		fmt.Println("Hola desde la funcion anonima")
	}()*/

	/*miFuncion := func(nombre string) string {
		return fmt.Sprintf("Hola, %s desde la funcion anonima", nombre)
	}
	fmt.Println(miFuncion("Thomas"))*/

}
