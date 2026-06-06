package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var jumlah float64
	var pi1, pi2 float64
	var i int

	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		suku := 1.0 / float64(2*i-1)

		if i%2 == 0 {
			suku = -suku
		}

		jumlah += suku

		// PI saat ini
		pi1 = 4 * jumlah

		// Hitung suku berikutnya
		sukuBerikut := 1.0 / float64(2*(i+1)-1)
		if (i+1)%2 == 0 {
			sukuBerikut = -sukuBerikut
		}

		// PI jika ditambah satu suku lagi
		pi2 = 4 * (jumlah + sukuBerikut)

		// Cek selisih dua pendekatan PI
		if math.Abs(pi2-pi1) < 0.00001 {
			break
		}
	}

	fmt.Printf("Hasil PI: %.9f\n", pi1)
	fmt.Printf("Hasil PI: %.9f\n", pi2)
	fmt.Printf("Pada i ke: %d\n", i)
}
