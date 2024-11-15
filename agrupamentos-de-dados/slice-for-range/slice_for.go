package main

import "fmt"

func main() {
	frutas := []string{"Grumixama", "Cereja do Rio Grande", "Ameixa Rubimel", "Laranja Champagne"}

	frutas = append(frutas, "Nectarina", "Maçã Julieta", "Maçã Eva", "Uva Vitória")

	for i, frutas := range frutas {
		fmt.Println("No índice", i,"temos a fruta:", frutas)
	}

}