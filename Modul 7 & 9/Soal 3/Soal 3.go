package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Scan(&klubA)
	fmt.Scan(&klubB)
	var skorA, skorB int
	pemenang := []string{}
	hasil := []string{}
	i := 1
	for {
		fmt.Scan(&skorA, &skorB)
		if skorA < 0 || skorB < 0 {
			break
		}
		if skorA > skorB {
			pemenang = append(pemenang, klubA)
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}
		i++
	}
	for i := 0; i < len(hasil); i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil[i])
	}
	fmt.Println("Pertandingan selesai")
}
