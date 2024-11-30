package main

import "fmt"

type Pessoa struct {
	Nome				string
	Sobrenome			string
	SaboresFavoritos 	[]string
}

func main() {
	pessoa1 := Pessoa{
		Nome: "Graça",
		Sobrenome: "Abreu",
		SaboresFavoritos: []string{"Chocolate", "Granola", "Morango", "Aveia"},
	}

	pessoa2 := Pessoa{
		Nome: "André",
		Sobrenome: "Abreu",
		SaboresFavoritos: []string{"Pitaia", "Chocolate", "Negresco", "Cajú"},
	}

	pessoas := []Pessoa{pessoa1, pessoa2}

	for _, pessoa := range pessoas {
		fmt.Printf("Nome: %s\n", pessoa.Nome)
		fmt.Println("Sabores favoritos:")
		for _, sabor := range pessoa.SaboresFavoritos {
			fmt.Printf("- %s\n", sabor)
		}
		fmt.Println()
	}

}