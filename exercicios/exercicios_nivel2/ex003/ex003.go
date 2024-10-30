package main

import "fmt"

const (
	ola = "Bom dia" // const não tipada
	oi string = "Bom dia, tudo bem?" // const tipada
)

func main() {
	fmt.Println(ola)
	fmt.Println(oi)
}