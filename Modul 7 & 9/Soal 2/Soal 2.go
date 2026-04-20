package main

import (
	"fmt"
	"math"
)

func tampilSemua(arr []int) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
func indeksGanjil(arr []int) {
	for i, v := range arr {
		if i%2 != 0 {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
}
func indeksGenap(arr []int) {
	for i, v := range arr {
		if i%2 == 0 {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
}
func kelipatanIndeks(arr []int, x int) {
	for i, v := range arr {
		if i%x == 0 {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
}
func hapusIndeks(arr []int, idx int) []int {
	return append(arr[:idx], arr[idx+1:]...)
}
func rataRata(arr []int) float64 {
	total := 0
	for _, v := range arr {
		total += v
	}
	return float64(total) / float64(len(arr))
}
func standarDeviasi(arr []int) float64 {
	mean := rataRata(arr)
	var jumlah float64
	for _, v := range arr {
		jumlah += math.Pow(float64(v)-mean, 2)
	}
	return math.Sqrt(jumlah / float64(len(arr)))
}
func frekuensi(arr []int, x int) int {
	count := 0
	for _, v := range arr {
		if v == x {
			count++
		}
	}
	return count
}
func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	tampilSemua(arr)
	indeksGanjil(arr)
	indeksGenap(arr)
	var x int
	fmt.Scan(&x)
	kelipatanIndeks(arr, x)
	var idx int
	fmt.Scan(&idx)
	arr = hapusIndeks(arr, idx)
	tampilSemua(arr)
	fmt.Println(rataRata(arr))
	fmt.Println(standarDeviasi(arr))
	var cari int
	fmt.Scan(&cari)
	fmt.Println(frekuensi(arr, cari))
}
