package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}
func didalam(cx, cy, r, x, y float64) bool {
	if jarak(cx, cy, x, y) <= r {
		return true
	}
	return false
}
func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x float64
	var y float64
	var dalam1 bool
	var dalam2 bool
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)
	dalam1 = didalam(cx1, cy1, r1, x, y)
	dalam2 = didalam(cx2, cy2, r2, x, y)
	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
