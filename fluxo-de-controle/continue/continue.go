package main

import "fmt"

func main() {
	for i := 0; i <= 20; i++ {
		if i % 2 != 0 {
			// pula para a próxima iteração
			// se o número for par, pula para a próxima iteração
			continue
		}
		fmt.Println(i)
	}
}