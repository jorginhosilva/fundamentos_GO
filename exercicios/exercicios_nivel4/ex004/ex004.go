package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	novoX := append(x, 52)

	fmt.Println(novoX)

	novoXX := append(x, 53, 54, 55)

	fmt.Println(novoXX)

	y := []int{56, 57, 58, 59, 60}

	novoXXX := append(x, y...)

	fmt.Println(novoXXX)

}