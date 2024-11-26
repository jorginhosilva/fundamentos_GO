package main

import "fmt"

func main() {
	estados := make([]string, 26)

	estados[0] = "Acre"
	estados[1] = "Alagoas"
	estados[2] = "Amapá"
	estados[3] = "Amazonas"
	estados[4] = "Bahia"
	estados[5] = "Ceará"
	estados[6] = "Espirito Santo"
	estados[7] = "Goiás"
	estados[8] = "Maranhão"
	estados[9] = "Mato Grosso"
	estados[10] = "Mato Grosso do Sul"
	estados[11] = "Minas Gerais"
	estados[12] = "Pará"
	estados[13] = "Paraíba"
	estados[14] = "Paraná"
	estados[15] = "Pernambuco"
	estados[16] = "Piauí"
	estados[17] = "Rio de Janeiro"
	estados[18] = "Rio Grande do Norte"
	estados[19] = "Rio Grande do Sul"
	estados[20] = "Rondônia"
	estados[21] = "Roraima"
	estados[22] = "Santa Catarina"
	estados[23] = "São Paulo"
	estados[24] = "Sergipe"
	estados[25] = "Tocantins"

	fmt.Println(len(estados))
	fmt.Println(cap(estados))
	
	for i := 0; i < len(estados); i++ {
		fmt.Println(estados[i])
	}

}