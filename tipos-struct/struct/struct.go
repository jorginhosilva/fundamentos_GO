package main

import "fmt"

// criando uma struct
// também uma forma de organizar seus dados
type Pessoa struct { // declara um novo tipo, chamado Pessoa.
		// Nome, Idade e Endereco -> são os campos da struct. Cada campo tem um tipo específico 
		Nome 		string
		Idade		int
		Endereco	string	
	}	

func main() {
	
	pessoa1 := Pessoa{"André", 28, "Rua Oton Gaspar de Farias, 47"}

	fmt.Println(pessoa1.Nome)
	fmt.Println(pessoa1.Idade)
	fmt.Println(pessoa1.Endereco)

}