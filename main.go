package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struktur Pengeluaran
// Kategori: transportasi, akomodasi, makanan, hiburan
// Jumlah: nilai pengeluaran

type Pengeluaran struct {
	Kategori string
	Jumlah   float64
}

var daftarPengeluaran []Pengeluaran
var anggaran float64

func main() {
	fmt.Println("=== Aplikasi Perencanaan Anggaran Perjalanan ===")
	anggaran = inputAngkaFloat("Masukkan total anggaran perjalanan: ")

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Pengeluaran")
		fmt.Println("2. Ubah Pengeluaran")
		fmt.Println("3. Hapus Pengeluaran")
		fmt.Println("4. Lihat Laporan")
		fmt.Println("5. Cari Pengeluaran (Sequential & Binary)")
		fmt.Println("6. Urutkan Pengeluaran (Selection & Insertion)")
		fmt.Println("7. Keluar")
		pilihan := inputAngkaInt("Pilih opsi: ")

		switch pilihan {
		case 1:
			tambahPengeluaran()
		case 2:
			ubahPengeluaran()
		case 3:
			hapusPengeluaran()
		case 4:
			cetakLaporan()
		case 5:
			cariPengeluaran()
		case 6:
			urutkanPengeluaran()
		case 7:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			os.Exit(0)
		default:
			fmt.Println("Opsi tidak valid.")
		}
	}
}

func inputAngkaFloat(prompt string) float64 {
	for {
		fmt.Print(prompt)
		var teks string
		fmt.Scanln(&teks)
		nilai, err := strconv.ParseFloat(strings.TrimSpace(teks), 64)
		if err != nil {
			fmt.Println("Input tidak valid, masukkan angka.")
			continue
		}
		return nilai
	}
}

func inputAngkaInt(prompt string) int {
	for {
		fmt.Print(prompt)
		var teks string
		fmt.Scanln(&teks)
		nilai, err := strconv.Atoi(strings.TrimSpace(teks))
		if err != nil {
			fmt.Println("Input tidak valid, masukkan angka integer.")
			continue
		}
		return nilai
	}
}

func inputTeks(prompt string) string {
	fmt.Print(prompt)
	var teks string
	fmt.Scanln(&teks)
	return strings.TrimSpace(teks)
}

func tambahPengeluaran() {
	kategori := inputTeks("Masukkan kategori (transportasi, akomodasi, makanan, hiburan): ")
	jumlah := inputAngkaFloat("Masukkan jumlah pengeluaran: ")

	daftarPengeluaran = append(daftarPengeluaran, Pengeluaran{Kategori: kategori, Jumlah: jumlah})
	fmt.Println("Pengeluaran berhasil ditambahkan.")
}

func ubahPengeluaran() {
	if len(daftarPengeluaran) == 0 {
		fmt.Println("Belum ada pengeluaran.")
		return
	}
	cetakSemuaPengeluaran()
	idx := inputAngkaInt("Pilih indeks pengeluaran yang akan diubah: ")
	if idx < 0 || idx >= len(daftarPengeluaran) {
		fmt.Println("Indeks tidak valid.")
		return
	}
	daftarPengeluaran[idx].Kategori = inputTeks("Masukkan kategori baru: ")
	daftarPengeluaran[idx].Jumlah = inputAngkaFloat("Masukkan jumlah baru: ")
	fmt.Println("Pengeluaran berhasil diubah.")
}

func hapusPengeluaran() {
	if len(daftarPengeluaran) == 0 {
		fmt.Println("Belum ada pengeluaran.")
		return
	}
	cetakSemuaPengeluaran()
	idx := inputAngkaInt("Pilih indeks pengeluaran yang akan dihapus: ")
	if idx < 0 || idx >= len(daftarPengeluaran) {
		fmt.Println("Indeks tidak valid.")
		return
	}

	daftarPengeluaran = append(daftarPengeluaran[:idx], daftarPengeluaran[idx+1:]...)
	fmt.Println("Pengeluaran berhasil dihapus.")
}

func cetakSemuaPengeluaran() {
	fmt.Println("Daftar Pengeluaran:")
	for i, p := range daftarPengeluaran {
		fmt.Printf("%d. %s: %.2f\n", i, p.Kategori, p.Jumlah)
	}
}

func totalPengeluaran() float64 {
	total := 0.0
	for _, p := range daftarPengeluaran {
		total += p.Jumlah
	}
	return total
}

func saranHemat() {
	dibelanjakan := totalPengeluaran()
	selisih := anggaran - dibelanjakan
	if selisih < 0 {
		fmt.Printf("Anda melebihi anggaran sebesar %.2f. Pertimbangkan untuk mengurangi pengeluaran.\n", -selisih)
	} else {
		fmt.Printf("Sisa anggaran Anda %.2f. Tetap hemat!\n", selisih)
	}
}

func cetakLaporan() {
	fmt.Println("=== Laporan Pengeluaran ===")
	perKategori := map[string]float64{}
	for _, p := range daftarPengeluaran { // _ untuk mengabaikan index
		perKategori[p.Kategori] += p.Jumlah
	}
	fmt.Println("Pengeluaran per kategori:")
	for kat, jml := range perKategori {
		fmt.Printf("- %s: %.2f\n", kat, jml)
	}
	fmt.Printf("Total pengeluaran: %.2f\n", totalPengeluaran())
	fmt.Printf("Anggaran: %.2f\n", anggaran)
	saranHemat()
}

func cariPengeluaran() {
	if len(daftarPengeluaran) == 0 {
		fmt.Println("Belum ada data untuk dicari.")
		return
	}
	query := strings.ToLower(inputTeks("Masukkan kategori yang dicari: "))

	// Binary Search melalui urutan insertion sort
	cats := make([]string, len(daftarPengeluaran)) // membuat slice untuk kategori dan membuatnya jadi lowercase
	for i, p := range daftarPengeluaran {
		cats[i] = strings.ToLower(p.Kategori)
	}
	// Insertion Sort pada cats (string di urut menggunakan urutan leksikografis)
	for i := 1; i < len(cats); i++ {
		kunci := cats[i]
		j := i - 1
		for j >= 0 && cats[j] > kunci {
			cats[j+1] = cats[j]
			j--
		}
		cats[j+1] = kunci
	}
	// Binary Search (O(log n))
	bawah, atas := 0, len(cats)-1
	ketemu := false
	for bawah <= atas {
		tengah := (bawah + atas) / 2
		if cats[tengah] == query {
			ketemu = true
			break
		} else if cats[tengah] < query {
			bawah = tengah + 1
		} else {
			atas = tengah - 1
		}
	}
	fmt.Println("Hasil Pencarian Binary:")
	if ketemu {
		fmt.Println("Kategori ditemukan pada daftar tersortir.")
	} else {
		fmt.Println("Kategori tidak ditemukan pada daftar tersortir.")
	}

	// Sequential Search (Linear Search) - O(n)
	hasilSeq := []Pengeluaran{}
	for _, p := range daftarPengeluaran {
		if strings.ToLower(p.Kategori) == query {
			hasilSeq = append(hasilSeq, p)
		}
	}
	fmt.Println("Hasil Pencarian Sequential:")
	if len(hasilSeq) == 0 {
		fmt.Println("Tidak ada pengeluaran pada kategori tersebut.")
	} else {
		for _, p := range hasilSeq {
			fmt.Printf("- %.2f\n", p.Jumlah)
		}
	}

}

func urutkanPengeluaran() {
	fmt.Println("Pilih metode:")
	fmt.Println("1. Berdasarkn Jumlah (Selection Sort)")
	fmt.Println("2. Berdasarkan Kategori (Insertion Sort)")
	metode := inputAngkaInt("Opsi: ")
	switch metode {
	case 1:
		selectionSortBerdasarkanJumlah()
		fmt.Println("Selesai Selection Sort berdasarkan jumlah.")
	case 2:
		insertionSortBerdasarkanKategori()
		fmt.Println("Selesai Insertion Sort berdasarkan kategori.")
	default:
		fmt.Println("Opsi tidak valid.")
	}
	cetakSemuaPengeluaran()
}

func selectionSortBerdasarkanJumlah() {
	n := len(daftarPengeluaran)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if daftarPengeluaran[j].Jumlah < daftarPengeluaran[minIdx].Jumlah {
				minIdx = j
			}
		}
		daftarPengeluaran[i], daftarPengeluaran[minIdx] = daftarPengeluaran[minIdx], daftarPengeluaran[i]
	}
}

func insertionSortBerdasarkanKategori() {
	n := len(daftarPengeluaran)
	for i := 1; i < n; i++ {
		kunci := daftarPengeluaran[i]
		j := i - 1
		for j >= 0 && strings.ToLower(daftarPengeluaran[j].Kategori) > strings.ToLower(kunci.Kategori) {
			daftarPengeluaran[j+1] = daftarPengeluaran[j]
			j--
		}
		daftarPengeluaran[j+1] = kunci
	}
}
