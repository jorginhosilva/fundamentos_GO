package main

import "fmt"

func main() {
	// podemos usar quantos ELSE IF quiser no código.
	// o else if só entra em ação, caso o IF seja negativo.
	// mas tem uma regrinha, tem que iniciar o código com um if e finalizar com um else.

	a := 25
	b := 26

	if (a < b) {
		fmt.Println("A é menor que B!")
	} else if (a > b) {
		fmt.Println("A é maior que B!")
	} else {
		fmt.Println("Ambos os números são iguais!")
	}

}