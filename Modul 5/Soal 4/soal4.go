package main

import "fmt"

func barisan(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	barisan(n - 1)
	if n != 1 {
		fmt.Print(n, " ")
	}
}
func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	barisan(n)
}
