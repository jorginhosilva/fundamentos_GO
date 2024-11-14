package main

import "fmt"

var diaDaSemana [7]string

func main() {
	
	diaDaSemana[0] = "Domingo"
	diaDaSemana[1] = "Segunda"
	diaDaSemana[2] = "Terça"
	diaDaSemana[3] = "Quarta"
	diaDaSemana[4] = "Quinta"
	diaDaSemana[5] = "Sexta"
	diaDaSemana[6] = "Sábado"

	fmt.Println(diaDaSemana)
	fmt.Println(len(diaDaSemana))

}