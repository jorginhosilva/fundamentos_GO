package main

import "fmt"

func main() {
	andre := make(map[string][]string)
	graca := make(map[string][]string)
	salete := make(map[string][]string)

	andre["ferreiraDasilva"] = []string{"Tomar cerveja", "Viajar pro Nordeste", "Estudar sobre Finanças"}
	graca["daPazabreu"] = []string{"Ler Livros", "Viajar pro Nordeste", "Estudar sobre o Direito"}
	salete["deJesus"] = []string{"Ir para a casa das filhas", "Ficar com os netos", "Ir pra fazenda 3 coqueiros"}

	fmt.Println(andre)
	fmt.Println(graca)
	fmt.Println(salete)

}