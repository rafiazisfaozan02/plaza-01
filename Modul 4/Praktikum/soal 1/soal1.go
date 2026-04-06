package main

import "fmt"

func faktorial(n int, hasil *int) {
	var i int
	*hasil = 1
	for i = 1; i <= n; i++ {
		*hasil = *hasil * i
	}
}
func permutasi(n, r int, hasil *int) {
	var fn, fnr int
	faktorial(n, &fn)
	faktorial(n-r, &fnr)
	*hasil = fn / fnr
}
func kombinasi(n, r int, hasil *int) {
	var fn, fr, fnr int
	faktorial(n, &fn)
	faktorial(r, &fr)
	faktorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}
func main() {
	var a, b, c, d int
	var p1, c1, p2, c2 int
	fmt.Scan(&a, &b, &c, &d)
	permutasi(a, c, &p1)
	kombinasi(a, c, &c1)
	permutasi(b, d, &p2)
	kombinasi(b, d, &c2)
	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
