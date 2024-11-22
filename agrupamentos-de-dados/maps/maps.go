package main

import "fmt"

func main() {
	// Criando um map para armazenar a idade de pessoas

	idades := make(map[string]int)

	idades["Alice"] = 30
	idades["Bob"] = 25
	idades["Laysa"] = 11
	idades["Luis Fernando"] = 8

	//idadeDeAlice := idades["Alice"]
	//fmt.Println(idadeDeAlice)

	for nome, idade := range idades {
		fmt.Printf("%s tem %d anos\n", nome, idade)
	}

	

}