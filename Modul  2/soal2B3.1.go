package main

import "fmt"

func main() {
	var kiri, kanan float64
	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri, &kanan)

		if kiri >= 9 || kanan >= 9 {
			break
		}
	}
	fmt.Println("Proses selesai.")
}
