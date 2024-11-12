package main

import "fmt"

func main() {
	
	for num := 10; num <= 100; num++ {
		resto := num % 4
		fmt.Printf("%d %% 4 = %d\n", num, resto)
	}


}