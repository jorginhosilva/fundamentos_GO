package main

import "fmt"

func main() {
	slice := []int{20,5,37,65,99,7,10,11,22,33}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}
	fmt.Printf("%T", slice)
}