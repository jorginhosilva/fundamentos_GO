package main

import "fmt"

func main() {
	// Criando um map para armazenar a idade de pessoas

	idades := make(map[string]int)

	idades["Alice"] = 30
	idades["Bob"] = 25

	idadeDeAlice := idades["Alice"]
	fmt.Println(idadeDeAlice)

}