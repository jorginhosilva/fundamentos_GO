package main

import "fmt"

func main() {
	// int -> São números inteiros, 10, 20, 30
    // float -> São números fracionarios, 10.0, 20.0, 30.0
    // em GO int e int32 são bem diferentes, o que difere entre eles é o tamanho do número suportado
    // usar somente float64

    a := 10 // número inteiro
    b := 10.0 // número float

    fmt.Printf("%v, %T\n", a, a)
    fmt.Printf("%v, %T\n", b, b)

}