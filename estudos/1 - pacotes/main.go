// main.go

package main

import (
	"fmt"

	"github.com/Vittor-Emanoel/Golang/estudos/1 - pacotes/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever(msg)("olá")
}
