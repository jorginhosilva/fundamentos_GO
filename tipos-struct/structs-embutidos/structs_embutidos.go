package main

import "fmt"

type Endereco struct {
    Rua     string
    Numero  int
    Cidade  string
    Estado  string
    CEP     int
}

type Pessoa struct {
    Nome        string
    Idade       int
    Endereco    Endereco
}

func main() {
	p1 := Pessoa{
        Nome: "André Luiz",
        Idade: 28,
        Endereco: Endereco{
            Rua: "Oton Gaspar de Farias",
            Numero: 44,
            Cidade: "Carneiros",
            Estado: "AL",
            CEP: 57535000,
        },
    }

    p2 := Pessoa{
        Nome: "Almir",
        Idade: 31,
        Endereco: Endereco{
            Rua: "Oton Gaspar de Farias",
            Numero: 44,
            Cidade: "Carneiros",
            Estado: "AL",
            CEP: 57535000,
        },
    }

    fmt.Println(p1.Nome, p1.Idade, p1.Endereco.Cidade,)
    fmt.Println(p2.Nome, p2.Idade, p2.Endereco.Cidade)

}