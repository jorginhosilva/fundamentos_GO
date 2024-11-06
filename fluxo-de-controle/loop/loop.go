package main

import "fmt"

func main() {
	// Não existe while em GO
	// for é o laço de repetição
	// x := 0 - declaramos uma variável com o nome x, e inicializamos ela com valor 0
	// x < 10 - é a condição de parada do laço
	// x++ - após cada execução do bloco de código dentro do laço, o valor de x será incrementado em 1
	// fmt.Println(x) - função responsável por imprimir algo na tela, nesse caso a variável de nome x
	
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
}