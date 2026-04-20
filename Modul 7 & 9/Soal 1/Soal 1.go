package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}
type Lingkaran struct {
	pusat Titik
	r     int
}

func jarak(a, b Titik) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	return math.Sqrt(dx*dx + dy*dy)
}
func diDalamLingkaran(t Titik, l Lingkaran) bool {
	return jarak(t, l.pusat) <= float64(l.r)
}
func main() {
	var l1, l2 Lingkaran
	var t Titik
	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&t.x, &t.y)
	dalam1 := diDalamLingkaran(t, l1)
	dalam2 := diDalamLingkaran(t, l2)
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
