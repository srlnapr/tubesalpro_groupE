package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX int = 100

type Pasien struct {
	ID     string
	Nama   string
	Umur   int
	Alamat string
}
type Layanan struct {
	ID    string
	Nama  string
	Harga int
}

type Kunjungan struct {
	IDKunjungan string
	IDPasien    string
	IDLayanan   string
	Tanggal     string
	Catatan     string
}

type Transaksi struct {
	IDTransaksi  string
	IDPasien     string
	NamaPasien   string
	Layanan      string
	TotalBiaya   float64
	TanggalVisit string
}

var dataKunjungan [MAX]Kunjungan
var jumlahKunjungan int
var dataPasien [MAX]Pasien
var jumlahPasien int
var dataLayanan [MAX]Layanan
var jumlahLayanan int
var dataTransaksi [MAX]Transaksi
var jumlahTransaksi int

var scanner = bufio.NewScanner(os.Stdin)

func tambahPasien() {
	var p Pasien

	fmt.Println("\n=== Tambah Pasien ===")

	fmt.Print("ID Pasien: ")
	fmt.Scan(&p.ID)

	if p.ID == "" {
		fmt.Println("ID tidak boleh kosong!")
		return
	}

	fmt.Print("Nama: ")
	fmt.Scan(&p.Nama)

	fmt.Print("Umur: ")
	fmt.Scan(&p.Umur)

	fmt.Print("Alamat: ")
	fmt.Scan(&p.Alamat)

	dataPasien[jumlahPasien] = p
	jumlahPasien++

	fmt.Println("Data pasien berhasil ditambahkan.")
}

func lihatPasien() {
	fmt.Println("\n=== Data Pasien ===")

	if jumlahPasien == 0 {
		fmt.Println("Blm ada data.")
		return
	}

	for i := 0; i < jumlahPasien; i++ {
		fmt.Println("Data ke-", i+1)
		fmt.Println("ID     :", dataPasien[i].ID)
		fmt.Println("Nama   :", dataPasien[i].Nama)
		fmt.Println("Umur   :", dataPasien[i].Umur)
		fmt.Println("Alamat :", dataPasien[i].Alamat)
		fmt.Println()
	}
}

func tambahLayanan() {
	var l Layanan

	fmt.Println("\n=== Tambah Layanan ===")

	fmt.Print("ID Layanan: ")
	fmt.Scan(&l.ID)

	if l.ID == "" {
		fmt.Println("ID tidak boleh kosong!")
		return
	}

	fmt.Print("Nama Layanan: ")
	fmt.Scan(&l.Nama)

	fmt.Print("Harga: ")
	fmt.Scan(&l.Harga)

	dataLayanan[jumlahLayanan] = l
	jumlahLayanan++

	fmt.Println("Data layanan berhasil ditambahkan.")
}

func lihatLayanan() {
	fmt.Println("\n=== Data Layanan ===")

	if jumlahLayanan == 0 {
		fmt.Println("Blm ada data.")
		return
	}

	for i := 0; i < jumlahLayanan; i++ {
		fmt.Println("Data ke-", i+1)
		fmt.Println("ID    :", dataLayanan[i].ID)
		fmt.Println("Nama  :", dataLayanan[i].Nama)
		fmt.Println("Harga :", dataLayanan[i].Harga)
		fmt.Println()
	}
}
func sequentialSearchPasien(data [MAX]Pasien, n int, keyword string) int {
	var i int = 0
	var idx int = -1
	var ketemu bool = false

	for i < n && !ketemu {
		if data[i].ID == keyword || data[i].Nama == keyword {
			idx = i
			ketemu = true
		}
		i++
	}

	return idx
}

func binarySearchPasienByID(data [MAX]Pasien, n int, keyword string) int {
	var kiri, kanan, tengah int
	var idx int = -1
	var ketemu bool = false

	kiri = 0
	kanan = n - 1

	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2

		if data[tengah].ID == keyword {
			idx = tengah
			ketemu = true
		} else if data[tengah].ID < keyword {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return idx
}

func catatKunjungan(data *[MAX]Kunjungan, n *int, daftarPasien [MAX]Pasien, jumlahPasien int) {
	var k Kunjungan
	var idx int

	fmt.Println("\n=== Catat Kunjungan Pasien ===")

	fmt.Print("ID Kunjungan: ")
	fmt.Scan(&k.IDKunjungan)

	fmt.Print("ID Pasien: ")
	fmt.Scan(&k.IDPasien)

	idx = sequentialSearchPasien(daftarPasien, jumlahPasien, k.IDPasien)

	if idx == -1 {
		fmt.Println("Pasien tidak ditemukan.")
		return
	}

	fmt.Print("ID Layanan: ")
	fmt.Scan(&k.IDLayanan)

	fmt.Print("Tanggal: ")
	fmt.Scan(&k.Tanggal)

	fmt.Print("Catatan Perawatan: ")
	fmt.Scan(&k.Catatan)

	data[*n] = k
	*n++

	fmt.Println("Kunjungan berhasil dicatat.")
}

func lihatKunjungan(data [MAX]Kunjungan, n int) {
	var i int

	fmt.Println("\n=== Riwayat Kunjungan ===")

	if n == 0 {
		fmt.Println("Belum ada data kunjungan.")
		return
	}

	for i = 0; i < n; i++ {
		fmt.Println("Data ke-", i+1)
		fmt.Println("ID Kunjungan :", data[i].IDKunjungan)
		fmt.Println("ID Pasien    :", data[i].IDPasien)
		fmt.Println("ID Layanan   :", data[i].IDLayanan)
		fmt.Println("Tanggal      :", data[i].Tanggal)
		fmt.Println("Catatan      :", data[i].Catatan)
		fmt.Println()
	}
}
func sortPasienByIDAsc(data *[MAX]Pasien, n int) {
	var i, j int
	var temp Pasien

	for i = 1; i < n; i++ {
		temp = data[i]
		j = i - 1

		for j >= 0 && data[j].ID > temp.ID {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = temp
	}
}
func cariPasienMenu(data [MAX]Pasien, n int) {
	var keyword string
	var metode int
	var idx int

	fmt.Println("\n=== Cari Pasien ===")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search berdasarkan ID")
	fmt.Print("Pilih metode: ")
	fmt.Scan(&metode)

	fmt.Print("Masukkan nama atau ID pasien: ")
	fmt.Scan(&keyword)

	if metode == 1 {
		idx = sequentialSearchPasien(data, n, keyword)
	} else if metode == 2 {
		sortPasienByIDAsc(&dataPasien, jumlahPasien)
		idx = binarySearchPasienByID(dataPasien, jumlahPasien, keyword)
	} else {
		idx = -1
		fmt.Println("Metode tidak valid.")
	}

	if idx != -1 {
		fmt.Println("Pasien ditemukan:")
		fmt.Println("ID     :", dataPasien[idx].ID)
		fmt.Println("Nama   :", dataPasien[idx].Nama)
		fmt.Println("Umur   :", dataPasien[idx].Umur)
		fmt.Println("Alamat :", dataPasien[idx].Alamat)
	} else {
		fmt.Println("Pasien tidak ditemukan.")
	}
}

func tambahTransaksi() {
	if jumlahTransaksi >= MAX {
		fmt.Println("Kapasitas penyimpanan transaksi penuh.")
		return
	}

	var t Transaksi
	fmt.Println("\n--- Tambah Transaksi Pembayaran ---")

	fmt.Print("ID Transaksi: ")
	scanner.Scan()
	t.IDTransaksi = strings.TrimSpace(scanner.Text())

	fmt.Print("ID Pasien: ")
	scanner.Scan()
	t.IDPasien = strings.TrimSpace(scanner.Text())

	fmt.Print("Nama Pasien: ")
	scanner.Scan()
	t.NamaPasien = strings.TrimSpace(scanner.Text())

	fmt.Print("Layanan: ")
	scanner.Scan()
	t.Layanan = strings.TrimSpace(scanner.Text())

	fmt.Print("Total Biaya: ")
	scanner.Scan()
	biayaStr := strings.TrimSpace(scanner.Text())
	t.TotalBiaya, _ = strconv.ParseFloat(biayaStr, 64)

	fmt.Print("Tanggal (YYYY-MM-DD): ")
	scanner.Scan()
	t.TanggalVisit = strings.TrimSpace(scanner.Text())

	dataTransaksi[jumlahTransaksi] = t
	jumlahTransaksi++
	fmt.Println("Transaksi berhasil ditambahkan.")

}

func tampilkanSemuaTransaksi(data [MAX]Transaksi, n int) {
	if n == 0 {
		fmt.Println(" (Belum ada data transaksi)")
		return
	}

	fmt.Printf("\n%-14s %-10s %-20s %-20s %12s %s\n",
		"ID Transaksi", "IDPasien", "Nama Paien", "Layanan", "Total Biaya", "Tanggal")
	fmt.Println(strings.Repeat("_", 95))

	for i := 0; i < n; i++ {
		fmt.Printf("%-14s %-10s %-20s %-20s %12.2f %s\n",
			data[i].IDTransaksi, data[i].IDPasien, data[i].NamaPasien, data[i].Layanan, data[i].TotalBiaya, data[i].TanggalVisit)

	}
}

func menuKelolaTransaksi() {
	for {
		fmt.Println("\n=== Kelola Transaksi Pembayaran ===")
		fmt.Println("1. Tambah Transaksi")
		fmt.Println("2. Lihat Semua Transaksi")
		fmt.Println("3. Hapus Transaksi")
		fmt.Println("0. Kembali")
		fmt.Print("Pilihan: ")

		var pilih int
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			tambahTransaksi()
		case 2:
			fmt.Println("\n--- Daftar Transaksi ---")
			tampilkanSemuaTransaksi(dataTransaksi, jumlahTransaksi)
		case 3:
			hapusTransaksi()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func hapusTransaksi() {
	fmt.Printf("\nMasukkan ID Transaksi yang akan dihapus: ")
	var id string
	fmt.Scanln(&id)

	for i := 0; i < jumlahTransaksi; i++ {
		if dataTransaksi[i].IDTransaksi == id {
			for j := i; j < jumlahTransaksi-1; j++ {
				dataTransaksi[j] = dataTransaksi[j+1]
			}

			jumlahTransaksi--
			fmt.Println("Transaksi berhasil dihapus.")
			return
		}
	}
	fmt.Println("ID Transaksi tidak ditemukan.")
}

func selectionSortTanggal(data [MAX]Transaksi, n int) [MAX]Transaksi {
	hasil := data

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if hasil[j].TanggalVisit < hasil[minIdx].TanggalVisit {
				minIdx = j
			}
		}
		hasil[i], hasil[minIdx] = hasil[minIdx], hasil[i]
	}
	return hasil
}

func insertionSortBiaya(data [MAX]Transaksi, n int) [MAX]Transaksi {
	hasil := data

	for i := 1; i < n; i++ {
		key := hasil[i]
		j := i - 1

		for j >= 0 && hasil[j].TotalBiaya < key.TotalBiaya {
			hasil[j+1] = hasil[j]
			j--
		}
		hasil[j+1] = key
	}
	return hasil
}

func menuUrutkanTransaksi() {
	if jumlahTransaksi == 0 {
		fmt.Println("\nBelum ada data transaksi untuk diurutkan.")
		return
	}

	fmt.Println("\n=== Urutkan Data Transaksi ===")
	fmt.Println("1. Urutkan berdasarkan Tanggal (Selection Sort - Ascending)")
	fmt.Println("2. Urutkan berdasarkan Total Biaya (Insertion Sort - Descending)")
	fmt.Println("0. Kembali")
	fmt.Print("Pilihan: ")

	var pilih int
	fmt.Scanln(&pilih)

	switch pilih {
	case 1:
		hasil := selectionSortTanggal(dataTransaksi, jumlahTransaksi)
		fmt.Println("\n--- Hasil Pengurutan berdasarkan Tanggal (Terlama -> Terbaru) ---")
		tampilkanSemuaTransaksi(hasil, jumlahTransaksi)
	case 2:
		hasil := insertionSortBiaya(dataTransaksi, jumlahTransaksi)
		fmt.Println("\n--- Hasil Pengurutan berdasarkan Total Biaya (Termalah -> Termurah) ---")
		tampilkanSemuaTransaksi(hasil, jumlahTransaksi)
	case 0:
		return
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func tampilkanStatistik() {
	if jumlahTransaksi == 0 {
		fmt.Println("\nBelum ada data transaksi.")
		return
	}

	fmt.Println("\n=== Statistik ===")

	kunjunganHarian := make(map[string]int)
	for i := 0; i < jumlahTransaksi; i++ {
		kunjunganHarian[dataTransaksi[i].TanggalVisit]++
	}

	fmt.Println("\nJumlah Kunjungan Harian:")
	fmt.Printf(" %-15s %s\n", "Tanggal", "Jumlah Kunjungan")
	fmt.Println(" " + strings.Repeat("-", 35))
	for tgl, jumlah := range kunjunganHarian {
		fmt.Printf(" %-15s %d\n", tgl, jumlah)
	}

	hitungLayanan := make(map[string]int)
	for i := 0; i < jumlahTransaksi; i++ {
		hitungLayanan[dataTransaksi[i].Layanan]++
	}

	layananTerbanyak := ""
	maxCount := 0
	for layanan, count := range hitungLayanan {
		if count > maxCount {
			maxCount = count
			layananTerbanyak = layanan
		}
	}

	fmt.Println("\nLayanan Paling Banyak Diminati:")
	fmt.Printf(" %-20s %s\n", "Layanan", "Jumlah Peminat")
	fmt.Println(" " + strings.Repeat("-", 35))

	for layanan, count := range hitungLayanan {
		tanda := ""
		if layanan == layananTerbanyak {
			tanda = " [terbanyak]"
		}
		fmt.Printf(" %-20s %d%s\n", layanan, count, tanda)
	}
}

func main() {
	var pilihan int
	var sisaEnter string

	for {
		fmt.Println("\n======= SIM-KLIK =======")
		fmt.Println("1. Tambah Pasien")
		fmt.Println("2. Lihat Pasien")
		fmt.Println("3. Tambah Layanan")
		fmt.Println("4. Lihat Layanan")
		fmt.Println("5. Catat Kunjungan")
		fmt.Println("6. Lihat Kunjungan")
		fmt.Println("7. Cari Pasien")
		fmt.Println("8. Kelola Transaksi Pembayaran")
		fmt.Println("9. Urutkan Data Transaksi")
		fmt.Println("10. Tampilkan Statistik")
		fmt.Println("11. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahPasien()
			fmt.Scanln(&sisaEnter)

		case 2:
			lihatPasien()
		case 3:
			tambahLayanan()
			fmt.Scanln(&sisaEnter)
		case 4:
			lihatLayanan()
		case 5:
			catatKunjungan(&dataKunjungan, &jumlahKunjungan, dataPasien, jumlahPasien)
			fmt.Scanln(&sisaEnter)
		case 6:
			lihatKunjungan(dataKunjungan, jumlahKunjungan)
		case 7:
			cariPasienMenu(dataPasien, jumlahPasien)
			fmt.Scanln(&sisaEnter)
		case 8:
			menuKelolaTransaksi()
		case 9:
			menuUrutkanTransaksi()
		case 10:
			tampilkanStatistik()
		case 11:
			fmt.Println("\nProgram selesai.")
			return
		default:
			fmt.Println("Menu tidak valid.")
			pilihan = 0
		}
	}
}
