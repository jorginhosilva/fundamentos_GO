package main

import "fmt"

func main() {

	valorArray := [5]int{10, 15, 20, 25, 30}

	for i, valor := range valorArray {
		fmt.Println(i, valor)
	}

	fmt.Printf("%T", valorArray)

}