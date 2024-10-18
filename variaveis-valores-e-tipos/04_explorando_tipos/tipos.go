package main

import "fmt"

var y int64 
var z string 

func main() {
	// Tipos em GO são estáticos, uma vez definido ele não terá mudança de tipo
	// Ao declarar uma variável para conter valores de um certo tipo, essa variável só poderar conter esse tipo
	// O tipo pode ser deduzido pelo compilador

	x := 10
	fmt.Println(x)

	y = 11
	fmt.Println(y)
	z = "isso é uma string"
	fmt.Println(z)
}