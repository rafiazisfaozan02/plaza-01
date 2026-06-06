package main

import "fmt"

func main() {
	var sum float64
	var count int
	fmt.Println("Masukkan bilangan real (ketik 9999 untuk mengakhiri):")
	for {
		var num float64
		fmt.Scan(&num)
		if num == 9999 {
			break
		}
		sum += num
		count++
	}
	if count == 0 {
		fmt.Println("Tidak ada bilangan yang dimasukkan.")
	} else {
		rerata := sum / float64(count)
		fmt.Printf("Jumlah bilangan : %d\n", count)
		fmt.Printf("Rerata          : %.2f\n", rerata)
	}
}
