package main

import "fmt"

func main() {
	// uma string é imútavel 
	// uma string é um slice de bytes (fatia de byte)

	s := "Oi, sou Developer"
	fmt.Printf("%v\n%T\n", s, s)

	sb := []byte(s)
	fmt.Printf("%v\n%T\n", sb, sb)
}