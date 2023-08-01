package main

import (
	"errors"
	"fmt"
)

func main() {
	//inteiros
	var numero int64 = -12666666665
	fmt.Println(numero)

	var numero2 uint32 = 15454
	fmt.Println(numero2)

	//alias
	//Int32 - RUNE
	var numero3 rune = 1234
	fmt.Println(numero3)

	//BYTE - UNIT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal float32  = 123.45
	fmt.Println(numeroReal)

	var numeroReal1 float64  = 123.45
	fmt.Println(numeroReal1)

	numeroReal3 := 1211.511
	fmt.Println(numeroReal3)


	///

	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := "b"
	fmt.Println(char)

	//

	//valor 0 - todo dado tem o valor inicial =  0

	var texto int
	fmt.Println(texto)

	var booleano bool 
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}