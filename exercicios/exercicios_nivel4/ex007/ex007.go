package main

import "fmt"

func main() {
	dados := [][]string{
				[]string{
					"André Luiz",
					"Jorginho",
					"Tomar Cerveja",
				},
				[]string{
					"Graça",
					"Gracinha",
					"Ler livros",
				},
				[]string{
					"Antônio",
					"Antônio de Albino",
					"Cochilar na cadeira",
				},
	}

	for _, v := range dados {
		fmt.Println(v)
	}

	fmt.Println("\n\n")

	for _, v := range dados {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}
	}


}