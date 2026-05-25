package main

import "fmt"

func selectionSortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
func selectionSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		var odd []int
		var even []int
		for j := 0; j < m; j++ {
			var x int
			fmt.Scan(&x)
			if x%2 != 0 {
				odd = append(odd, x)
			} else {
				even = append(even, x)
			}
		}
		selectionSortAsc(odd)
		selectionSortDesc(even)
		first := true
		for _, v := range odd {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(v)
			first = false
		}
		for _, v := range even {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(v)
			first = false
		}
		fmt.Println()
	}
}
