package main

import "fmt"

func main() {
	// SWITCH compostas, é realizar comparações mais complexas.
	// esse tipo de switch não é muit comum.
	// para agrupar múltiplos valores em um único case, usa virgula ,.
	// FALLTHROUGH permite que a execução continue para o próximo caso.

	nota := 6

	switch nota {
	case 9, 10:
		fmt.Println("Excelente!")
	case 7, 8:
		fmt.Println("Bom")
	default:
		fmt.Println("Precisa melhorar!")
	}


}