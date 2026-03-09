package main

import "fmt"

func main() {
	var b int
	var prima bool
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)
	fmt.Print("Faktor: ")
	countFaktor := 0
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			countFaktor++
		}
	}
	fmt.Println()
	if countFaktor == 2 {
		prima = true
	} else if countFaktor != 2 {
		prima = false
	}
	fmt.Println("Prima: ", prima)
}
