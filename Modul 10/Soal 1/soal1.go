package main

import "fmt"

func main() {
	var N int
	var berat [1000]float64
	var min float64
	var max float64
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)
	if N <= 0 || N > 1000 {
		fmt.Println("Jumlah harus antara 1 sampai 1000")
		return
	}
	for i := 0; i < N; i++ {
		fmt.Scan(&berat[i])
	}
	min = berat[0]
	max = berat[0]
	for i := 1; i < N; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}
	fmt.Println("Berat terkecil: ", min)
	fmt.Println("Berat terbesar: ", max)
}
