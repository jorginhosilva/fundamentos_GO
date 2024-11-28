package main

import "fmt"

func main() {
	numbers := []int{10,20,30,40,50,60,70,80,90,100}

	fmt.Println(numbers)

	fatia_numbers := numbers[0:3]

	fmt.Println(fatia_numbers)

	fatia_numbers2 := numbers[4:]

	fmt.Println(fatia_numbers2)

	fatia_numbers3 := numbers[1:8]

	fmt.Println(fatia_numbers3)

	fatia_numbers4 := numbers[2:9]

	fmt.Println(fatia_numbers4)

	fmt.Println(numbers[2: len(numbers)-1])

}