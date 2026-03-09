package main

import "fmt"

func main() {
	var parsel int
	var sisa int
	var biaya int
	var kg int
	fmt.Print("Masukkan berat parsel: ")
	fmt.Scan(&parsel)
	sisa = parsel % 1000
	kg = parsel / 1000
	if parsel < 10000 && sisa >= 500 {
		biaya = (parsel-sisa)*10 + sisa*5
		fmt.Println("Detail berat: ", kg, "kg +", sisa, "gr")
		fmt.Println("Total biaya: Rp.", biaya)
	} else if parsel < 10000 && sisa < 500 {
		biaya = (parsel-sisa)*10 + sisa*15
		fmt.Println("Detail berat: ", kg, "kg +", sisa, "gr")
		fmt.Println("Total biaya: Rp.", biaya)
	} else if parsel > 10000 && sisa < 1000 {
		biaya = (parsel - sisa) * 10
		fmt.Println("Detail berat: ", kg, "kg +", sisa, "gr")
		fmt.Println("Total biaya: Rp.", biaya)
	}
}
