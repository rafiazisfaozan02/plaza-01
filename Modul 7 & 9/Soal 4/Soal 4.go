package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	for {
		fmt.Scanf("%c", &ch)
		if ch == '.' || *n >= NMAX {
			break
		}
		t[*n] = ch
		*n++
	}
}
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}
func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}
func main() {
	var tab tabel
	var n int
	isiArray(&tab, &n)
	cetakArray(tab, n)
	if palindrome(tab, n) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan palindrome")
	}
	balikanArray(&tab, n)
	cetakArray(tab, n)
}
