package main

import "fmt"

type veiculo struct {
	portas		int
	cor 		string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro	bool
}

type sedan struct {
	veiculo
	modeloLuxo	bool
}

func main() {
	carroPasseio := sedan{veiculo{4, "Branco"}, true}
	carroRoca := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor: "Branco",
		},
		tracaoNasQuatro: true,
	}

	fmt.Println(carroPasseio)
	fmt.Println(carroRoca)
	fmt.Println(carroPasseio.cor)
	fmt.Println(carroRoca.tracaoNasQuatro)

}