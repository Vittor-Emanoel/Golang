package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {
	somar := somar(10, 20)
	fmt.Println(somar)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Olá teste")
}
