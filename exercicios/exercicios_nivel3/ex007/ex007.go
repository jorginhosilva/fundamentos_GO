package main

import "fmt"

func main() {
	idade := 28

	if idade < 18 {
		fmt.Println("Menor de idade!")
	} else if idade >= 18 && idade < 60 {
		fmt.Println("Adulto!")
	} else {
		fmt.Println("Idoso!")
	}

}