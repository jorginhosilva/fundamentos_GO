package main

import "fmt"

func main() {
	esporteFavorito := "futsal"

	switch esporteFavorito{
	case "futebol":
		fmt.Println("Eba! você prefere futebol!")
	case "futsal":
		fmt.Println("Eba! você prefere futsal!")
	case "volei":
		fmt.Println("Eba! você prefere volei")
	default:
		fmt.Println("Esse esporte não está na lista de opções ou não existe.")
	}

}