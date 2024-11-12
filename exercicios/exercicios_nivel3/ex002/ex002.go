package main

import (
	"fmt"
)

func main() {
	for i := 'A'; i <= 'Z'; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("O código Unicode de '%c' é U+%X\n", i, i)
		}
		fmt.Println()
	}
}