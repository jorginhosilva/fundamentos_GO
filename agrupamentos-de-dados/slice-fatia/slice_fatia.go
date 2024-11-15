package main

import "fmt"

func main() {
	numeros := []int{5, 7, 12, 20, 33, 55, 72 }

	fmt.Println(numeros)

	fatia := numeros[1:5]

	fmt.Println(fatia)

	// o novo slice nao copia os elementos do slice original
	// ele simplesmente cria um novo slice que aponta para a 
	// mesma região de memória do slice original
	// mas com um intervalo diferente

	num_fatia := numeros[:]

	fmt.Println(num_fatia)

}
