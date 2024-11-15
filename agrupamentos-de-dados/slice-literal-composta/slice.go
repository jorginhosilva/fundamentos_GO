package main

import "fmt"

func main() {
  // Criando um slice
  slice := []int{10, 20, 30}

	fmt.Println(slice)

  // Adicionando elementos
  slice = append(slice, 40, 50)

	fmt.Println(slice)

}