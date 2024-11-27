package main

import "fmt"

func main() {
	hobbies := map[string][]string{
		"daSilva_andre": []string{
			"tomar vários tipos de cervejas", "Viajar pro nordeste", "Estudar sobre Finanças",
		},
		"daPazabreu_graca": []string{
			"Ler livros", "Viajar pro nordeste", "Estudar sobre o Direito",
		},
		"deJesus_salete": []string{
			"Ir para a casa das filhas", "Ficar com os netos em casa", "Ir pro sítio de pai",
		},
	}

	hobbies["ferreira_antonio"] = []string{"Ir pro seu sitio", "Dormir na cadeira", "Assistir jogo no campo"}

	for k, v := range hobbies {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}

}