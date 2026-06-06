package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	fmt.Print("Masukkan banyaknya tetesan air hujan: ")
	fmt.Scan(&n)
	countA, countB, countC, countD := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if x < 0.5 && y < 0.5 {
			countA++
		} else if x >= 0.5 && y < 0.5 {
			countB++
		} else if x >= 0.5 && y >= 0.5 {
			countC++
		} else {
			countD++
		}
	}
	const perTetes = 0.0001
	fmt.Printf("Curah hujan daerah A: %.4f milimeter\n", float64(countA)*perTetes)
	fmt.Printf("Curah hujan daerah B: %.4f milimeter\n", float64(countB)*perTetes)
	fmt.Printf("Curah hujan daerah C: %.4f milimeter\n", float64(countC)*perTetes)
	fmt.Printf("Curah hujan daerah D: %.4f milimeter\n", float64(countD)*perTetes)
}
