package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MesinKarakter struct {
	untaian string
	posisi  int
}

var mesin MesinKarakter

func start(input string) {
	mesin.untaian = input
	mesin.posisi = 0
}

func maju() {
	mesin.posisi++
}

func eop() bool {
	return mesin.posisi >= len(mesin.untaian) || mesin.untaian[mesin.posisi] == '.'
}

func cc() byte {
	return mesin.untaian[mesin.posisi]
}

func bacaSemuaKarakter() {
	fmt.Print("Karakter terbaca: ")
	for !eop() {
		fmt.Printf("%c ", cc())
		maju()
	}
	fmt.Println()
}

func hitungKarakter() int {
	jumlah := 0
	for !eop() {
		jumlah++
		maju()
	}
	return jumlah
}

func hitungHurufA() int {
	jumlah := 0
	for !eop() {
		if cc() == 'A' {
			jumlah++
		}
		maju()
	}
	return jumlah
}

func hitungFrekuensiA() float64 {
	jumlahA := 0
	total := 0
	for !eop() {
		if cc() == 'A' {
			jumlahA++
		}
		total++
		maju()
	}
	if total == 0 {
		return 0
	}
	return float64(jumlahA) / float64(total) * 100
}

func hitungKataLE() int {
	jumlah := 0
	for !eop() {
		if cc() == 'L' {
			maju()
			if !eop() && cc() == 'E' {
				jumlah++
			}
		} else {
			maju()
		}
	}
	return jumlah
}

func tampilkanMenu() {
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║        MESIN ABSTRAK - KARAKTER              ║")
	fmt.Println("╠══════════════════════════════════════════════╣")
	fmt.Println("║  1. Baca Semua Karakter                      ║")
	fmt.Println("║  2. Hitung Jumlah Karakter                   ║")
	fmt.Println("║  3. Hitung Huruf 'A'                         ║")
	fmt.Println("║  4. Hitung Frekuensi Huruf 'A' (%)           ║")
	fmt.Println("║  5. Hitung Kata 'LE'                         ║")
	fmt.Println("║  6. Ganti Input                              ║")
	fmt.Println("║  0. Keluar                                   ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Print("Pilih menu: ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("==============================================")
	fmt.Println("     SELAMAT DATANG DI MESIN KARAKTER       ")
	fmt.Println("==============================================")
	fmt.Println("Catatan: Input harus diakhiri dengan titik (.)")
	fmt.Println("Contoh : HALO DUNIA BELAJAR GOLANG.")
	fmt.Print("\nMasukkan untaian karakter: ")
	scanner.Scan()
	input := strings.ToUpper(scanner.Text())

	if !strings.HasSuffix(input, ".") {
		input += "."
	}

	start(input)
	fmt.Printf("Input diterima: %s\n", input)

	for {
		tampilkanMenu()
		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {

		case "1":
			fmt.Println("\n--- Baca Semua Karakter ---")
			start(input)
			bacaSemuaKarakter()

		case "2":
			fmt.Println("\n--- Hitung Jumlah Karakter ---")
			start(input)
			jumlah := hitungKarakter()
			fmt.Printf("Jumlah karakter (tidak termasuk titik): %d\n", jumlah)

		case "3":
			fmt.Println("\n--- Hitung Huruf 'A' ---")
			start(input)
			jumlahA := hitungHurufA()
			fmt.Printf("Jumlah huruf 'A': %d\n", jumlahA)

		case "4":
			fmt.Println("\n--- Frekuensi Huruf 'A' ---")
			start(input)
			frekuensi := hitungFrekuensiA()
			fmt.Printf("Frekuensi huruf 'A': %.2f%%\n", frekuensi)

		case "5":
			fmt.Println("\n--- Hitung Kata 'LE' ---")
			start(input)
			jumlahLE := hitungKataLE()
			fmt.Printf("Jumlah kata 'LE': %d\n", jumlahLE)

		case "6":
			fmt.Println("\n--- Ganti Input ---")
			fmt.Println("Catatan: Input harus diakhiri dengan titik (.)")
			fmt.Print("Masukkan untaian karakter baru: ")
			scanner.Scan()
			input = strings.ToUpper(scanner.Text())
			if !strings.HasSuffix(input, ".") {
				input += "."
			}
			start(input)
			fmt.Printf("Input baru diterima: %s\n", input)

		case "0":
			fmt.Println("\nTerima kasih! Program selesai.")
			return

		default:
			fmt.Println(">> Pilihan tidak valid, coba lagi.")
		}
	}
}
