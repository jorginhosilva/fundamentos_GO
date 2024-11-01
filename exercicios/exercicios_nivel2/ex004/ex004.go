package main

import "fmt"

func main() {

    num := 5
    y := num << 1
    z := y

    fmt.Println("Variável num!")
    fmt.Printf("%d\n %b\n %#X\n", num, num, num)
    fmt.Println("Variável z!")
    fmt.Printf("%d\n %b\n %#X\n", z, z, z)
}