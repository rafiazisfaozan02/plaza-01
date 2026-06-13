package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Domino struct {
	sisi1   int
	sisi2   int
	nilai   int
	isBalak bool
}

type Dominoes struct {
	kartu  [28]Domino
	jumlah int
}

func buatDominoes() Dominoes {
	var set Dominoes
	idx := 0
	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			set.kartu[idx] = Domino{
				sisi1:   i,
				sisi2:   j,
				nilai:   i + j,
				isBalak: i == j,
			}
			idx++
		}
	}
	set.jumlah = 28
	return set
}

func kocokKartu(set *Dominoes) {
	rand.Seed(time.Now().UnixNano())
	for i := set.jumlah - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		set.kartu[i], set.kartu[j] = set.kartu[j], set.kartu[i]
	}
	fmt.Println(">> Kartu berhasil dikocok!")
}

func ambilKartu(set *Dominoes) Domino {
	if set.jumlah == 0 {
		fmt.Println(">> Tumpukan kartu sudah habis!")
		return Domino{}
	}
	set.jumlah--
	kartu := set.kartu[set.jumlah]
	return kartu
}

func gambarKartu(kartu Domino, suit int) int {
	if kartu.sisi1 == suit {
		return kartu.sisi2
	} else if kartu.sisi2 == suit {
		return kartu.sisi1
	}
	return -1
}

func nilaiKartu(kartu Domino) int {
	return kartu.nilai
}

func galiKartu(set *Dominoes, kartuAcuan Domino) Domino {
	fmt.Printf("\nMencari kartu dengan suit %d atau %d...\n", kartuAcuan.sisi1, kartuAcuan.sisi2)
	for set.jumlah > 0 {
		kartuAmbil := ambilKartu(set)
		fmt.Printf("  Mengambil kartu (%d|%d)... ", kartuAmbil.sisi1, kartuAmbil.sisi2)
		if kartuAmbil.sisi1 == kartuAcuan.sisi1 ||
			kartuAmbil.sisi1 == kartuAcuan.sisi2 ||
			kartuAmbil.sisi2 == kartuAcuan.sisi1 ||
			kartuAmbil.sisi2 == kartuAcuan.sisi2 {
			fmt.Println("COCOK!")
			return kartuAmbil
		}
		fmt.Println("tidak cocok.")
	}
	fmt.Println(">> Tidak ditemukan kartu yang cocok, tumpukan habis!")
	return Domino{}
}

func sepasangKartu(kartu1 Domino, kartu2 Domino) bool {
	return (kartu1.nilai + kartu2.nilai) == 12
}

func cetakKartu(kartu Domino) {
	balak := ""
	if kartu.isBalak {
		balak = " [BALAK]"
	}
	fmt.Printf("  Kartu: (%d|%d) | Nilai: %d%s\n", kartu.sisi1, kartu.sisi2, kartu.nilai, balak)
}

func tampilkanMenu() {
	fmt.Println("\n╔══════════════════════════════════════╗")
	fmt.Println("║       MESIN ABSTRAK - DOMINO         ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║  1. Kocok Kartu                      ║")
	fmt.Println("║  2. Ambil Kartu                      ║")
	fmt.Println("║  3. Gambar Kartu (cek nilai suit)    ║")
	fmt.Println("║  4. Nilai Kartu                      ║")
	fmt.Println("║  5. Lihat Semua Kartu di Tumpukan    ║")
	fmt.Println("║  6. Lihat Kartu di Tangan            ║")
	fmt.Println("║  7. Gali Kartu                       ║")
	fmt.Println("║  8. Sepasang Kartu (total = 12?)     ║")
	fmt.Println("║  0. Keluar                           ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Print("Pilih menu: ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	set := buatDominoes()
	var tangan []Domino

	fmt.Println("======================================")
	fmt.Println("   SELAMAT DATANG DI MESIN DOMINO    ")
	fmt.Println("======================================")
	fmt.Printf("Set kartu siap: %d kartu\n", set.jumlah)

	for {
		tampilkanMenu()
		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {

		case "1":
			fmt.Println("\n--- Kocok Kartu ---")
			kocokKartu(&set)
			fmt.Printf("Jumlah kartu di tumpukan: %d\n", set.jumlah)

		case "2":
			fmt.Println("\n--- Ambil Kartu ---")
			if set.jumlah == 0 {
				fmt.Println(">> Tumpukan kosong, tidak bisa ambil kartu!")
			} else {
				kartu := ambilKartu(&set)
				tangan = append(tangan, kartu)
				fmt.Print(">> Kartu yang diambil: ")
				cetakKartu(kartu)
				fmt.Printf("Sisa kartu di tumpukan: %d\n", set.jumlah)
			}

		case "3":
			fmt.Println("\n--- Gambar Kartu ---")
			if len(tangan) == 0 {
				fmt.Println(">> Kamu belum punya kartu! Ambil kartu dulu (menu 2).")
				break
			}
			fmt.Println("Kartu di tanganmu:")
			for i, k := range tangan {
				fmt.Printf("  [%d] ", i+1)
				cetakKartu(k)
			}
			fmt.Print("Pilih nomor kartu: ")
			scanner.Scan()
			var idxKartu int
			fmt.Sscan(scanner.Text(), &idxKartu)
			if idxKartu < 1 || idxKartu > len(tangan) {
				fmt.Println(">> Nomor kartu tidak valid!")
				break
			}
			kartuDipilih := tangan[idxKartu-1]
			fmt.Printf("Masukkan suit (0-6) untuk kartu (%d|%d): ", kartuDipilih.sisi1, kartuDipilih.sisi2)
			scanner.Scan()
			var suit int
			fmt.Sscan(scanner.Text(), &suit)
			hasil := gambarKartu(kartuDipilih, suit)
			if hasil == -1 {
				fmt.Printf(">> Suit %d tidak ada pada kartu (%d|%d)!\n", suit, kartuDipilih.sisi1, kartuDipilih.sisi2)
			} else {
				fmt.Printf(">> Kartu (%d|%d) dengan suit=%d -> nilai sisi lainnya = %d\n",
					kartuDipilih.sisi1, kartuDipilih.sisi2, suit, hasil)
			}

		case "4":
			fmt.Println("\n--- Nilai Kartu ---")
			if len(tangan) == 0 {
				fmt.Println(">> Kamu belum punya kartu! Ambil kartu dulu (menu 2).")
				break
			}
			fmt.Println("Kartu di tanganmu:")
			for i, k := range tangan {
				fmt.Printf("  [%d] ", i+1)
				cetakKartu(k)
			}
			fmt.Print("Pilih nomor kartu: ")
			scanner.Scan()
			var idxKartu int
			fmt.Sscan(scanner.Text(), &idxKartu)
			if idxKartu < 1 || idxKartu > len(tangan) {
				fmt.Println(">> Nomor kartu tidak valid!")
				break
			}
			kartuDipilih := tangan[idxKartu-1]
			fmt.Printf(">> Nilai kartu (%d|%d) = %d\n",
				kartuDipilih.sisi1, kartuDipilih.sisi2, nilaiKartu(kartuDipilih))

		case "5":
			fmt.Println("\n--- Kartu di Tumpukan ---")
			if set.jumlah == 0 {
				fmt.Println(">> Tumpukan kosong!")
			} else {
				fmt.Printf("Total: %d kartu\n", set.jumlah)
				for i := 0; i < set.jumlah; i++ {
					fmt.Printf("  [%d] ", i+1)
					cetakKartu(set.kartu[i])
				}
			}

		case "6":
			fmt.Println("\n--- Kartu di Tangan ---")
			if len(tangan) == 0 {
				fmt.Println(">> Kamu belum punya kartu!")
			} else {
				fmt.Printf("Total: %d kartu\n", len(tangan))
				for i, k := range tangan {
					fmt.Printf("  [%d] ", i+1)
					cetakKartu(k)
				}
			}

		case "7":
			fmt.Println("\n--- Gali Kartu ---")
			if len(tangan) == 0 {
				fmt.Println(">> Kamu belum punya kartu acuan! Ambil kartu dulu (menu 2).")
				break
			}
			if set.jumlah == 0 {
				fmt.Println(">> Tumpukan kosong, tidak bisa menggali!")
				break
			}
			fmt.Println("Pilih kartu acuan dari tanganmu:")
			for i, k := range tangan {
				fmt.Printf("  [%d] ", i+1)
				cetakKartu(k)
			}
			fmt.Print("Pilih nomor kartu acuan: ")
			scanner.Scan()
			var idxAcuan int
			fmt.Sscan(scanner.Text(), &idxAcuan)
			if idxAcuan < 1 || idxAcuan > len(tangan) {
				fmt.Println(">> Nomor kartu tidak valid!")
				break
			}
			kartuAcuan := tangan[idxAcuan-1]
			fmt.Printf("Kartu acuan: (%d|%d)\n", kartuAcuan.sisi1, kartuAcuan.sisi2)
			kartuHasil := galiKartu(&set, kartuAcuan)
			if kartuHasil.nilai != 0 || kartuHasil.isBalak {
				fmt.Print(">> Kartu ditemukan: ")
				cetakKartu(kartuHasil)
				tangan = append(tangan, kartuHasil)
			}
			fmt.Printf("Sisa kartu di tumpukan: %d\n", set.jumlah)

		case "8":
			fmt.Println("\n--- Sepasang Kartu ---")
			if len(tangan) < 2 {
				fmt.Println(">> Kamu butuh minimal 2 kartu di tangan!")
				break
			}
			fmt.Println("Kartu di tanganmu:")
			for i, k := range tangan {
				fmt.Printf("  [%d] ", i+1)
				cetakKartu(k)
			}
			fmt.Print("Pilih kartu pertama: ")
			scanner.Scan()
			var idx1 int
			fmt.Sscan(scanner.Text(), &idx1)
			fmt.Print("Pilih kartu kedua: ")
			scanner.Scan()
			var idx2 int
			fmt.Sscan(scanner.Text(), &idx2)
			if idx1 < 1 || idx1 > len(tangan) || idx2 < 1 || idx2 > len(tangan) || idx1 == idx2 {
				fmt.Println(">> Nomor kartu tidak valid!")
				break
			}
			k1 := tangan[idx1-1]
			k2 := tangan[idx2-1]
			fmt.Printf("\nKartu 1: (%d|%d) nilai=%d\n", k1.sisi1, k1.sisi2, k1.nilai)
			fmt.Printf("Kartu 2: (%d|%d) nilai=%d\n", k2.sisi1, k2.sisi2, k2.nilai)
			fmt.Printf("Total nilai: %d\n", k1.nilai+k2.nilai)
			if sepasangKartu(k1, k2) {
				fmt.Println(">> Hasil: TRUE - Total nilai kedua kartu adalah 12!")
			} else {
				fmt.Println(">> Hasil: FALSE - Total nilai kedua kartu bukan 12.")
			}

		case "0":
			fmt.Println("\nTerima kasih! Program selesai.")
			return

		default:
			fmt.Println(">> Pilihan tidak valid, coba lagi.")
		}
	}
}
