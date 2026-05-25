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
func checkJarak(arr []int, n int) {
	if n <= 1 {
		fmt.Println("Data berjarak tidak tetap")
		return
	}
	jarak := arr[1] - arr[0]
	tetap := true
	for i := 2; i < n; i++ {
		if arr[i]-arr[i-1] != jarak {
			tetap = false
			break
		}
	}
	if tetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
func main() {
	var data [1000]int
	count := 0
	var x int
	for {
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		data[count] = x
		count++
	}
	insertionSort(data[:], count)
	for i := 0; i < count; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(data[i])
	}
	fmt.Println()
	checkJarak(data[:count], count)
}
