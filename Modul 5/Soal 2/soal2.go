package main

import "fmt"

func cetakBintang(n int, j int) {
	if j > n {
		return
	}
	for i := 0; i < j; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	cetakBintang(n, j+1)
}
func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	cetakBintang(n, 1)
}
