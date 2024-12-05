package main

import "fmt"

func soma(a int, b int) int {
	return a + b
	// func -> palavra-chave que indica o início da definição de uma função.
	// soma -> Nome que você escolhe para identificar a função.
	// (a int, b int) -> parâmetros que a função recebe, com seus respectivos tipos.
	// int -> tipo de retorno que a função retorna. Se não houver retorno, pode ser omitido.
	// corpo da função -> bloco de código que contém as instruções a serem executadas pela função.
}

func main() {
	resultado := soma(3, 5)
	fmt.Println(resultado)
}