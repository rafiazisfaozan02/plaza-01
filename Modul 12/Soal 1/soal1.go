package main

import (
	"fmt"
	"sort"
)

func main() {
	votes := []int{}
	validVotes := map[int]int{}
	totalInput := 0
	for {
		var num int
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		totalInput++
		if num >= 1 && num <= 20 {
			validVotes[num]++
		}
	}
	totalValid := 0
	for _, count := range validVotes {
		totalValid += count
	}
	for candidate := range validVotes {
		votes = append(votes, candidate)
	}
	sort.Ints(votes)
	fmt.Printf("Suara masuk: %d\n", totalInput)
	fmt.Printf("Suara sah: %d\n", totalValid)
	for _, candidate := range votes {
		fmt.Printf("%d: %d\n", candidate, validVotes[candidate])
	}
}
