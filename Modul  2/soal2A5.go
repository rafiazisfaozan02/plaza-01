package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var x, y, z int
	fmt.Scan(&a, &b, &c, &d, &e)
	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)
	fmt.Scanf("%c%c%c", &x, &y, &z)
	fmt.Printf("%c%c%c\n", x+3, y+3, z+3)
}
