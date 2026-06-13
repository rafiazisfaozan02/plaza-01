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

type Pemain struct {
	nama   string
	tangan []Domino
}

var scanner = bufio.NewScanner(os.Stdin)

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
}

func ambilKartu(set *Dominoes) Domino {
	if set.jumlah == 0 {
		return Domino{sisi1: -1, sisi2: -1}
	}
	set.jumlah--
	return set.kartu[set.jumlah]
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
	for set.jumlah > 0 {
		kartuAmbil := ambilKartu(set)
		if kartuAmbil.sisi1 == kartuAcuan.sisi1 ||
			kartuAmbil.sisi1 == kartuAcuan.sisi2 ||
			kartuAmbil.sisi2 == kartuAcuan.sisi1 ||
			kartuAmbil.sisi2 == kartuAcuan.sisi2 {
			return kartuAmbil
		}
	}
	return Domino{sisi1: -1, sisi2: -1}
}

func cetakKartu(kartu Domino) string {
	return fmt.Sprintf("(%d|%d)", kartu.sisi1, kartu.sisi2)
}

func cetakRangkaian(rangkaian []Domino) {
	if len(rangkaian) == 0 {
		fmt.Println("  [kosong]")
		return
	}
	for i, k := range rangkaian {
		if i == len(rangkaian)-1 {
			fmt.Printf("%s", cetakKartu(k))
		} else {
			fmt.Printf("%s -- ", cetakKartu(k))
		}
	}
	fmt.Println()
}

func ujungKiri(rangkaian []Domino) int {
	return rangkaian[0].sisi1
}

func ujungKanan(rangkaian []Domino) int {
	return rangkaian[len(rangkaian)-1].sisi2
}

func bisaSambung(kartu Domino, rangkaian []Domino) bool {
	if len(rangkaian) == 0 {
		return true
	}
	kiri := ujungKiri(rangkaian)
	kanan := ujungKanan(rangkaian)
	return kartu.sisi1 == kiri || kartu.sisi2 == kiri ||
		kartu.sisi1 == kanan || kartu.sisi2 == kanan
}

func sambungKiri(kartu Domino, rangkaian []Domino) []Domino {
	kiri := ujungKiri(rangkaian)
	if kartu.sisi2 == kiri {
		return append([]Domino{kartu}, rangkaian...)
	}
	kartuBalik := Domino{sisi1: kartu.sisi2, sisi2: kartu.sisi1, nilai: kartu.nilai, isBalak: kartu.isBalak}
	return append([]Domino{kartuBalik}, rangkaian...)
}

func sambungKanan(kartu Domino, rangkaian []Domino) []Domino {
	kanan := ujungKanan(rangkaian)
	if kartu.sisi1 == kanan {
		return append(rangkaian, kartu)
	}
	kartuBalik := Domino{sisi1: kartu.sisi2, sisi2: kartu.sisi1, nilai: kartu.nilai, isBalak: kartu.isBalak}
	return append(rangkaian, kartuBalik)
}

func bagikanKartu(set *Dominoes, pemain []Pemain, jumlahKartuPerPemain int) {
	for i := range pemain {
		for j := 0; j < jumlahKartuPerPemain; j++ {
			kartu := ambilKartu(set)
			pemain[i].tangan = append(pemain[i].tangan, kartu)
		}
	}
}

func tampilkanTangan(pemain Pemain) {
	fmt.Printf("Kartu %s:\n", pemain.nama)
	for i, k := range pemain.tangan {
		fmt.Printf("  [%d] %s nilai=%d", i+1, cetakKartu(k), k.nilai)
		if k.isBalak {
			fmt.Print(" [BALAK]")
		}
		fmt.Println()
	}
}

func hapusKartuDariTangan(pemain *Pemain, idx int) {
	pemain.tangan = append(pemain.tangan[:idx], pemain.tangan[idx+1:]...)
}

func giliranPemain(pemain *Pemain, set *Dominoes, rangkaian *[]Domino) bool {
	fmt.Printf("\n========== Giliran %s ==========\n", pemain.nama)
	fmt.Printf("Sisa kartu di tumpukan: %d\n", set.jumlah)

	if len(*rangkaian) > 0 {
		fmt.Printf("Ujung rangkaian: KIRI=[%d] | KANAN=[%d]\n", ujungKiri(*rangkaian), ujungKanan(*rangkaian))
	}

	fmt.Print("Rangkaian saat ini: ")
	cetakRangkaian(*rangkaian)
	tampilkanTangan(*pemain)

	punyaKartuCocok := false
	for _, k := range pemain.tangan {
		if bisaSambung(k, *rangkaian) {
			punyaKartuCocok = true
			break
		}
	}

	if !punyaKartuCocok {
		if set.jumlah > 0 {
			fmt.Printf("%s tidak punya kartu yang cocok, mengambil kartu dari tumpukan...\n", pemain.nama)
			kartuBaru := ambilKartu(set)
			pemain.tangan = append(pemain.tangan, kartuBaru)
			fmt.Printf("Kartu diambil: %s\n", cetakKartu(kartuBaru))
			if !bisaSambung(kartuBaru, *rangkaian) {
				fmt.Printf("%s masih tidak bisa menyambung, giliran dilewat.\n", pemain.nama)
				return false
			}
		} else {
			fmt.Printf("%s tidak bisa menyambung dan tumpukan habis, giliran dilewat.\n", pemain.nama)
			return false
		}
	}

	fmt.Print("Pilih nomor kartu yang ingin disambung (0 = lewat): ")
	scanner.Scan()
	var pilihanKartu int
	fmt.Sscan(scanner.Text(), &pilihanKartu)

	if pilihanKartu == 0 {
		fmt.Printf("%s memilih lewat.\n", pemain.nama)
		return false
	}

	if pilihanKartu < 1 || pilihanKartu > len(pemain.tangan) {
		fmt.Println(">> Nomor kartu tidak valid, giliran dilewat.")
		return false
	}

	kartuDipilih := pemain.tangan[pilihanKartu-1]

	if !bisaSambung(kartuDipilih, *rangkaian) {
		fmt.Println(">> Kartu tidak bisa disambung, giliran dilewat.")
		return false
	}

	if len(*rangkaian) == 0 {
		*rangkaian = append(*rangkaian, kartuDipilih)
		hapusKartuDariTangan(pemain, pilihanKartu-1)
		fmt.Printf("Kartu %s diletakkan sebagai kartu pertama.\n", cetakKartu(kartuDipilih))
		return true
	}

	kiri := ujungKiri(*rangkaian)
	kanan := ujungKanan(*rangkaian)
	bisaKiri := kartuDipilih.sisi1 == kiri || kartuDipilih.sisi2 == kiri
	bisaKanan := kartuDipilih.sisi1 == kanan || kartuDipilih.sisi2 == kanan

	posisi := ""
	if bisaKiri && bisaKanan {
		fmt.Print("Sambung ke KIRI atau KANAN? (K/N): ")
		scanner.Scan()
		posisi = scanner.Text()
	} else if bisaKiri {
		posisi = "K"
	} else {
		posisi = "N"
	}

	if posisi == "K" || posisi == "k" {
		*rangkaian = sambungKiri(kartuDipilih, *rangkaian)
		fmt.Printf("Kartu %s disambung ke KIRI.\n", cetakKartu(kartuDipilih))
	} else {
		*rangkaian = sambungKanan(kartuDipilih, *rangkaian)
		fmt.Printf("Kartu %s disambung ke KANAN.\n", cetakKartu(kartuDipilih))
	}

	hapusKartuDariTangan(pemain, pilihanKartu-1)
	return true
}

func main() {
	fmt.Println("==============================================")
	fmt.Println("       PERMAINAN GAPLEH - DOMINO            ")
	fmt.Println("==============================================")

	fmt.Print("Masukkan jumlah pemain (2-4): ")
	scanner.Scan()
	var jumlahPemain int
	fmt.Sscan(scanner.Text(), &jumlahPemain)
	if jumlahPemain < 2 || jumlahPemain > 4 {
		fmt.Println("Jumlah pemain tidak valid, diset ke 2.")
		jumlahPemain = 2
	}

	pemain := make([]Pemain, jumlahPemain)
	for i := range pemain {
		fmt.Printf("Masukkan nama pemain %d: ", i+1)
		scanner.Scan()
		pemain[i].nama = scanner.Text()
	}

	jumlahKartuPerPemain := 28 / jumlahPemain
	if jumlahKartuPerPemain > 7 {
		jumlahKartuPerPemain = 7
	}

	set := buatDominoes()
	kocokKartu(&set)
	bagikanKartu(&set, pemain, jumlahKartuPerPemain)

	fmt.Println("\nKartu sudah dibagikan!")
	fmt.Printf("Setiap pemain mendapat %d kartu\n", jumlahKartuPerPemain)

	var rangkaian []Domino
	giliranSaat := 0
	maxRonde := 100

	for ronde := 0; ronde < maxRonde; ronde++ {
		p := &pemain[giliranSaat]
		giliranPemain(p, &set, &rangkaian)

		if len(p.tangan) == 0 {
			fmt.Println("\n==============================================")
			fmt.Printf("  SELAMAT! %s MENANG!\n", p.nama)
			fmt.Println("==============================================")
			fmt.Print("\nRangkaian akhir: ")
			cetakRangkaian(rangkaian)
			return
		}

		giliranSaat = (giliranSaat + 1) % jumlahPemain
	}

	fmt.Println("\n==============================================")
	fmt.Println("  PERMAINAN SELESAI! Tidak ada pemenang.")
	fmt.Println("==============================================")
	for _, p := range pemain {
		total := 0
		for _, k := range p.tangan {
			total += nilaiKartu(k)
		}
		fmt.Printf("  %s: %d kartu tersisa, total nilai = %d\n", p.nama, len(p.tangan), total)
	}
}
