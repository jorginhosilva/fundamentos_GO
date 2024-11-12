package main

import "fmt"

func main() {

	festa := 3

	fmt.Println("Qual data comemorativa você prefere?")
	fmt.Println("Digite 1 para carnaval:")
	fmt.Println("Digite 2 para São João:")
	fmt.Println("Digite 3 para nenhuma das opções:")

	switch {
	case festa == 1:
		fmt.Println("Eba!! você escolheu carnaval!")
	case festa == 2:
		fmt.Println("Eba!! você escolheu São João!")
	case festa == 3:
		fmt.Println("Você é caseiro, não curte festas.")
	}

}