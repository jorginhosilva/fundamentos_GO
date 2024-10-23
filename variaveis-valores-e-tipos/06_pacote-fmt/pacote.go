package main

import "fmt"

func main() {
	// Usando o Sprintln()
	nome := "André Luiz"
	idade := 28

	fmt.Sprintln("Sou o ,", nome, "! minha idade é: ", idade, "anos.")
	// Usando o Println()
	nome2 := "André Luiz"
	idade2 := 28

	fmt.Println("Meu nome é, ", nome2, "! minha idade é, ", idade2, "anos.")
}