package main

import (
	"fmt"
	"sort"
)

func main() {
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
	candidates := []int{}
	for candidate := range validVotes {
		candidates = append(candidates, candidate)
	}
	sort.Ints(candidates)
	maxVote1 := -1
	ketua := -1
	for _, c := range candidates {
		if validVotes[c] > maxVote1 {
			maxVote1 = validVotes[c]
			ketua = c
		}
	}
	maxVote2 := -1
	wakil := -1
	for _, c := range candidates {
		if c == ketua {
			continue
		}
		if validVotes[c] > maxVote2 {
			maxVote2 = validVotes[c]
			wakil = c
		}
	}
	fmt.Printf("Suara masuk: %d\n", totalInput)
	fmt.Printf("Suara sah: %d\n", totalValid)
	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
