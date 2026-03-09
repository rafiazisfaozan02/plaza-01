package main

import "fmt"

func main() {
	var N, i int
	var bunga string
	var pita string
	fmt.Scan(&N)
	for i = 1; i <= N; i++ {
		fmt.Scan(&bunga)
		pita = pita + bunga + " - "
	}
	fmt.Println("Pita:", pita)
}
