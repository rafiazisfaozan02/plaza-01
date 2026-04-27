package main

import "fmt"

func main() {
	var x, y int
	var ikan [1000]float64
	fmt.Scan(&x, &y)
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}
	var totalWadah []float64
	for i := 0; i < x; i += y {
		var sum float64 = 0
		for j := i; j < i+y && j < x; j++ {
			sum += ikan[j]
		}
		totalWadah = append(totalWadah, sum)
	}
	var total float64 = 0
	for i := 0; i < len(totalWadah); i++ {
		fmt.Print(totalWadah[i], " ")
		total += totalWadah[i]
	}
	fmt.Println()
	rata := total / float64(len(totalWadah))
	fmt.Println(rata)
}
