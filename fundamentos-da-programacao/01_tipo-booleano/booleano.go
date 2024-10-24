package main

import "fmt"

var x bool

func main() {
	// O tipo bool ou booleano só pode conter dois valores true and false (Bool é um tipo binário)
	// verdadeiro ou falso, sim ou não, zero ou um

	fmt.Println(x) // zero value de um bool será sempre falso
	x = true
	fmt.Println(x)

	x = (10 < 100)

	y := (10 > 15)

	z := (10 == 10)

	fmt.Println(y)
	fmt.Println(z)

}