package main

import "fmt"

func main() {

	// Make é utilizada para alocar e inicializar tipos de dados que possuem um tamanho dinâmico

	numeros := make([]int, 5)
	fmt.Println(numeros)
	numeros[0], numeros[1], numeros[2], numeros[3], numeros[4] = 1, 2, 3, 4, 5 

	fmt.Println(numeros)

}