package main

import "fmt"

func main() {
	var w1, w2, w3, w4 string
	var hasil bool
	hasil = true
	for i := 0; i < 5; i++ {
		fmt.Scan(&w1, &w2, &w3, &w4)
		hasil = hasil && (w1 == "merah" && w2 == "kuning" && w3 == "hijau" && w4 == "ungu")
	}
	fmt.Println("BERHASIL:", hasil)
}
