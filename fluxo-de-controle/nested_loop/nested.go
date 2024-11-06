package main

import "fmt"

func main() {
	// nested significa aninhado, no contexto atual é loop aninhado
	// muito útil quando usamos em slices.

	for dia := 0; dia < 12; dia++ {
		fmt.Println("Hora:", dia, "\nBom dia!")
		for x := 0; x <= 60; x++ {
			fmt.Print(" ", x)
		}
		fmt.Println("\n")
	}

	for tarde := 12; tarde <= 18; tarde++ {
		fmt.Println("Hora:", tarde, "\nBoa tarde!")
		for x := 0; x <= 60; x++ {
			fmt.Print(" ", x)
		}
		fmt.Println("\n")
	} 

	for noite := 19; noite <= 24; noite++ {
		fmt.Println("Hora:", noite, "\nBoa noite!")
		for x := 0; x <= 60; x++ {
			fmt.Print(" ", x)
		}
		fmt.Println("\n")
	}

}