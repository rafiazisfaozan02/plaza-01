package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	isiArray(n)

	result := posisi(n, k)

	if result == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(result)
	}
}
func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}
func posisi(n, k int) int {
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		if data[mid] == k {
			return mid
		} else if data[mid] < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
