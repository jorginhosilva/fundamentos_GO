package main

import "fmt"

// Aqui é uma declaração, onde pedimos pro pc reservar a memória pra tal variável
var x int
var a int
var b float64
var c string
var d bool

func main() {
	// Valor zero, é o valor que se encontra antes da variável ser iniciallizada
	// ints: 0
	// float: 0.0
	// booleans: false
	// strings: ""
	// pointers, functions, interfaces, slices, channels, maps: nil
	
	// use := operador curto de declaração sempre que possível
	// use var para package-level scope

	// Aqui é uma inicialização e atribuição
	x = 10
	fmt.Printf("%v, %T\n", x, x)

	// Aqui é uma atribuição
	x = 20 
	fmt.Printf("%v, %T\n", x, x)

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

}