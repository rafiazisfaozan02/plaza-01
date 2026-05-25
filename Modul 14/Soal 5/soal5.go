package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}
type DaftarBuku [nMax]Buku

var Pustaka DaftarBuku
var nPustaka int

func DaftarkanBuku(n int) {
	nPustaka = n
	for i := 0; i < n; i++ {
		fmt.Scan(&Pustaka[i].id)
		fmt.Scan(&Pustaka[i].judul)
		fmt.Scan(&Pustaka[i].penulis)
		fmt.Scan(&Pustaka[i].penerbit)
		fmt.Scan(&Pustaka[i].eksemplar)
		fmt.Scan(&Pustaka[i].tahun)
		fmt.Scan(&Pustaka[i].rating)
	}
}
func UrutBuku() {
	for i := 1; i < nPustaka; i++ {
		key := Pustaka[i]
		j := i - 1
		for j >= 0 && Pustaka[j].rating < key.rating {
			Pustaka[j+1] = Pustaka[j]
			j--
		}
		Pustaka[j+1] = key
	}
}
func CetakTerfavorit() {
	maxIdx := 0
	for i := 1; i < nPustaka; i++ {
		if Pustaka[i].rating > Pustaka[maxIdx].rating {
			maxIdx = i
		}
	}
	b := Pustaka[maxIdx]
	fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\n",
		b.judul, b.penulis, b.penerbit, b.tahun)
}
func Cetak5Terbaru() {
	limit := 5
	if nPustaka < 5 {
		limit = nPustaka
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("%s\n", Pustaka[i].judul)
	}
}
func CariBuku(r int) {
	left := 0
	right := nPustaka - 1
	found := -1
	for left <= right {
		mid := (left + right) / 2
		if Pustaka[mid].rating == r {
			found = mid
			break
		} else if Pustaka[mid].rating > r {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if found == -1 {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	} else {
		b := Pustaka[found]
		fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nEksemplar: %d\nRating: %d\n",
			b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
	}
}
func main() {
	var n, r int
	fmt.Scan(&n)
	DaftarkanBuku(n)
	fmt.Println("=== Buku Terfavorit ===")
	CetakTerfavorit()
	UrutBuku()
	fmt.Println("\n=== 5 Buku Rating Tertinggi ===")
	Cetak5Terbaru()
	fmt.Scan(&r)
	fmt.Println("\n=== Hasil Pencarian ===")
	CariBuku(r)
}
