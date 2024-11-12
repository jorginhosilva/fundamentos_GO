package main

import "fmt"

func main() {
	nasci := 1995
	atual := 2024

	for {
		if nasci > atual {
			break
		}
		fmt.Println(nasci)
		nasci++
	}

}