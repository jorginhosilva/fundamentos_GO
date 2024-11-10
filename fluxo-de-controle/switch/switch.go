package main

import "fmt"

func main() {

	// SWITCH é basicamente um IF mais facil de trabalhar e mais conciso de entender o código.
	// por padrão não há fall-through.

	dia := 3

	switch dia {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda")
	case 3:
		fmt.Println("Terça")
	case 4: 
		fmt.Println("Quarta")
	default:
		fmt.Println("Dia Inválido!")
	}
	
}