package main

import "fmt"

func main() {
	var kiri, kanan float64
	var selisih float64
	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri, &kanan)
		if kiri < 0 || kanan < 0 || kiri+kanan > 150 {
			break
		}
		selisih = kiri - kanan
		if selisih < 0 {
			selisih = -selisih
		}
		fmt.Println("Sepeda motor pak Andi akan oleng:", selisih >= 9)
	}
	fmt.Println("Proses selesai.")
}
