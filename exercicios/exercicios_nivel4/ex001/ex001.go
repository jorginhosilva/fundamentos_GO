package main

import "fmt"

//var valor [5]int64

func main() {

	valorArray := [5]int{10, 15, 20, 25, 30}

	for i, valor := range valorArray {
		fmt.Println(i, valor)
	}

	fmt.Printf("%T", valorArray)

}