package main

import "fmt"

type Pessoa struct {
	Nome				string
	Sobrenome			string
	SaboresFavoritos 	[]string
}

func main() {
	mappessoas := make(map[string]Pessoa)

	mappessoas["Abreu"] = Pessoa{
		Nome: "Graça",
		Sobrenome: "Abreu",
		SaboresFavoritos: []string{"Chocolate", "Granola", "Morango", "Aveia"},
	}

	mappessoas["Ferreira"] = Pessoa{
		Nome: "André",
		Sobrenome: "Ferreira",
		SaboresFavoritos: []string{"Pitaia", "Chocolate", "Negresco", "Cajú"},
	}
	fmt.Println("\n")
	for _, valorpessoas := range mappessoas {
		fmt.Printf("Meu nome é: %s\n", valorpessoas.Nome)
		fmt.Printf("Os meus sorvetes favoritos são: %s\n", valorpessoas.SaboresFavoritos)
	}
	fmt.Println("\n")
}

// Demonstre os valores do map utilizando range.