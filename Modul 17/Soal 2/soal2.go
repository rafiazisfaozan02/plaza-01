package main

import "fmt"

func main() {
	var x string
	fmt.Print("Masukkan string x: ")
	fmt.Scan(&x)
	var n int
	fmt.Print("Masukkan jumlah data (n): ")
	fmt.Scan(&n)
	data := make([]string, n)
	fmt.Println("Masukkan", n, "buah string:")
	for i := 0; i < n; i++ {
		fmt.Printf("  Data ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}
	ditemukan := false
	for _, s := range data {
		if s == x {
			ditemukan = true
			break
		}
	}
	fmt.Println("\n=== HASIL ===")
	if ditemukan {
		fmt.Printf("a. String \"%s\" ADA dalam kumpulan data.\n", x)
	} else {
		fmt.Printf("a. String \"%s\" TIDAK ADA dalam kumpulan data.\n", x)
	}
	fmt.Printf("b. Posisi string \"%s\" ditemukan: ", x)
	posisi := []int{}
	for i, s := range data {
		if s == x {
			posisi = append(posisi, i+1)
		}
	}
	if len(posisi) == 0 {
		fmt.Println("tidak ditemukan.")
	} else {
		fmt.Println(posisi)
	}
	jumlah := len(posisi)
	fmt.Printf("c. Jumlah string \"%s\" dalam kumpulan data: %d\n", x, jumlah)
	if jumlah >= 2 {
		fmt.Printf("d. Terdapat SEDIKITNYA DUA string \"%s\" dalam kumpulan data.\n", x)
	} else {
		fmt.Printf("d. TIDAK terdapat sedikitnya dua string \"%s\" dalam kumpulan data.\n", x)
	}
}
