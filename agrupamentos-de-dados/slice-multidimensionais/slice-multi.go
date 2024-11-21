package main

import "fmt"

func main() {
	// Criando uma matriz 2x3

	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(matriz)

	valor := matriz[1][2]
	fmt.Println(valor)

}