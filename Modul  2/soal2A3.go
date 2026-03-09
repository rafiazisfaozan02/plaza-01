package main

import "fmt"

func main() {
	var r float64
	fmt.Print("Masukkan nilai jejari: ")
	fmt.Scan(&r)
	phi := 3.1415926535
	volume := (4.0 / 3.0) * phi * r * r * r
	luas := 4 * phi * r * r
	fmt.Printf("Bola dengan jejari %.0f memiliki Volume %.4f dan Luas %.4f", r, volume, luas)

}
