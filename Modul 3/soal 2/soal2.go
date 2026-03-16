package main

import "fmt"

func komposisif(x int) int {
	var fungsif int
	fungsif = x * x
	return fungsif
}
func komposisig(x int) int {
	var fungsig int
	fungsig = x - 2
	return fungsig
}
func komposisih(x int) int {
	var fungsih int
	fungsih = x + 1
	return fungsih
}
func main() {
	var a int
	var b int
	var c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(komposisif(komposisig(komposisih(a))))
	fmt.Println(komposisig(komposisih(komposisif(b))))
	fmt.Println(komposisih(komposisif(komposisig(c))))
}
