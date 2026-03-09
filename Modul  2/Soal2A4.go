package main

import "fmt"

func main() {
	var celsius float64
	fmt.Print("Masukkan nilai derajat celcius: ")
	fmt.Scan(&celsius)
	reamur := celsius * 4 / 5
	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273
	fmt.Println("Derajat Reamur:", reamur)
	fmt.Println("Derajat Fahrenheit:", fahrenheit)
	fmt.Println("Derajat Kelvin:", kelvin)

}
