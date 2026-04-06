package main

import "fmt"

func hitungSkor(jumlahSoal *int, totalWaktu *int) {
	var waktu int
	var i int
	*jumlahSoal = 0
	*totalWaktu = 0
	for i = 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 300 {
			*jumlahSoal++
			*totalWaktu += waktu
		}
	}
}
func main() {
	var nama string
	var pemenang string
	var maxSoal, minWaktu int
	maxSoal = -1
	minWaktu = 999999
	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}
		var soal, waktu int
		hitungSkor(&soal, &waktu)
		if soal > maxSoal || (soal == maxSoal && waktu < minWaktu) {
			maxSoal = soal
			minWaktu = waktu
			pemenang = nama
		}
	}
	fmt.Println(pemenang, maxSoal, minWaktu)
}
