package main

import "fmt"

func main() {

	// A função append é uma ferramenta para adicionar elementos a um slice.
	// Os ... são obrigatórios para que o código funcione.

	numeros := []int{1, 2, 3, 4, 5}
	outrosNumeros := []int{16, 17, 18, 19, 20}

	fmt.Println(numeros)

	numeros = append(numeros, 6, 7, 8, 9, 10)

	fmt.Println(numeros)

	numeros = append(numeros, outrosNumeros...)

	fmt.Println(numeros)


}