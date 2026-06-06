package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var n int
	fmt.Print("Banyak Topping: ")
	fmt.Scan(&n)
	xc, yc, r := 0.5, 0.5, 0.5
	count := 0
	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if math.Pow(x-xc, 2)+math.Pow(y-yc, 2) <= r*r {
			count++
		}
	}
	pi := 4.0 * float64(count) / float64(n)
	fmt.Printf("Topping pada Pizza: %d\n", count)
	fmt.Printf("PI : %.10f\n", pi)
}
