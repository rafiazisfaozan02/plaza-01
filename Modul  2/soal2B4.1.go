package main

import "fmt"

func main() {
	var k float64
	var f float64
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)
	f = (4*k + 2) * (4*k + 2) / ((4*k + 1) * (4*k + 3))
	fmt.Printf("Nilai f(K) = %.9f\n", f)
}
