package main

import "fmt"

func main() {

	// IF é uma condicional que sempre terá um valor bool, caso esse valor seja um not,
	// a condicional do IF não apresentará nenhum resultado

	x := 25 

	if (x > 20) {
		fmt.Println("Sim, x é maior!")
	}
}