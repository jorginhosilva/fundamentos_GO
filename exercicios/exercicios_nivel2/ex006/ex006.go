package main

import "fmt"

const (
	anoatual = 2024
	ano1 = anoatual + iota
	ano2 = anoatual + iota
	ano3 = anoatual + iota
	ano4 = anoatual + iota
)

func main() {
	fmt.Println(ano1)
	fmt.Println(ano2)
	fmt.Println(ano3)
	fmt.Println(ano4)
}