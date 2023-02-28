package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {
	valores := strings.Split(expresion, "+")
	var result int
	for _, valor := range valores {
		numb, error := strconv.Atoi(valor)
		if error != nil {
			fmt.Println(error)
			fmt.Println("Error: Tiene que ingresar un numero entero")
			fmt.Println("o !Solo debes realizar con el operador +!!")
		} else {
			result += numb
		}
	}

	return result
}

func main() {
	var expresion string
	var result int

	fmt.Print("=>")
	fmt.Scanln(&expresion)
	result = sumar(expresion)
	fmt.Printf("=> %d\n", result)
}
