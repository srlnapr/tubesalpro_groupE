package main

import "fmt"

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

var dataPasien [MAX]Pasien
var jumlahPasien int
var dataLayanan [MAX]Layanan
var jumlahLayanan int

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

func main() {
	var pilihan int

	for {
		fmt.Println("\n=== MODUL ORANG 1 ===")
		fmt.Println("1. Tambah Pasien")
		fmt.Println("2. Lihat Pasien")
		fmt.Println("3. Tambah Layanan")
		fmt.Println("4. Lihat Layanan")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahPasien()
		case 2:
			lihatPasien()
		case 3:
			tambahLayanan()
		case 4:
			lihatLayanan()
		case 5:
			fmt.Println("Program selesai.")
			return
		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}
