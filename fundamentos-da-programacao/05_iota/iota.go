package main

import "fmt"

const (
	x = iota
	y = iota
	z = iota
)

func main() {
	// iota funciona como um contador automático dentro de declarações const

	fmt.Println(x, y, z)

}