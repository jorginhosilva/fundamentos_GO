package main

import "fmt"

func main() {

	ano_nasci := 1995
	ano_atual := 2024
	
	for ano_nasci <= ano_atual {
		fmt.Println(ano_nasci)
		ano_nasci++
	}
}