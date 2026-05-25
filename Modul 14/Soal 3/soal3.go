package main

import "fmt"

func insertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
func getMedian(arr []int, n int) int {
	if n%2 == 0 {
		return (arr[n/2-1] + arr[n/2]) / 2
	}
	return arr[n/2]
}
func main() {
	var data [1000000]int
	count := 0
	var x int
	for {
		fmt.Scan(&x)
		if x == -5313 {
			break
		}
		if x == 0 {
			insertionSort(data[:], count)
			fmt.Println(getMedian(data[:count], count))
		} else {
			data[count] = x
			count++
		}
	}
}
