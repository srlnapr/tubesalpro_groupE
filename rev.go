package main

import (
	"fmt"
)

const NMAX = 100

type User struct {
	Username string
	Password string
	Role     string
}

type Pasien struct {
	ID     int
	Nama   string
	Alamat string
	NoHP   string
}

type Layanan struct {
	ID    int
	Nama  string
	Harga float64
}

type Transaksi struct {
	IDPasien   int
	NamaPasien string
	Layanan    string
	Tanggal    int
	TotalBiaya float64
}

type Kunjungan struct {
	IDPasien   int
	NamaPasien string
	Layanan    string
	Tanggal    int
}

// Variabel global hanya untuk array utama
var dataUser [NMAX]User
var dataPasien [NMAX]Pasien
var dataLayanan [NMAX]Layanan
var dataTransaksi [NMAX]Transaksi
var dataKunjungan [NMAX]Kunjungan

func banner() {
	fmt.Println("*****************************************************")
	fmt.Println("***      Aplikasi Manajemen Klinik (SIM-KLIK)      ***")
	fmt.Println("***          Algoritma Pemrograman  2025           ***")
	fmt.Println("*****************************************************")
}

func menuUtama() {
	fmt.Println("\n*** Menu Utama ***")
	fmt.Println("1. Registrasi")
	fmt.Println("2. Login")
	fmt.Println("3. Keluar")
	fmt.Println("----------------------")
}

// registrasi menambahkan user baru ke array dataUser
// Parameter: nUser *int - pointer ke jumlah user saat ini
func registrasi(nUser *int) {
	var pilih int

	if *nUser >= NMAX {
		fmt.Println("Data user penuh")
		return
	}

	fmt.Println("\n*** REGISTRASI ***")
	fmt.Print("Username : ")
	fmt.Scan(&dataUser[*nUser].Username)

	fmt.Print("Password : ")
	fmt.Scan(&dataUser[*nUser].Password)

	fmt.Println("Pilih Role :")
	fmt.Println("1. Admin")
	fmt.Println("2. Kasir")
	fmt.Print("Pilih role [1-2] : ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		dataUser[*nUser].Role = "admin"
	} else if pilih == 2 {
		dataUser[*nUser].Role = "kasir"
	} else {
		fmt.Println("Role tidak valid")
		return
	}

	(*nUser)++
	fmt.Println("Registrasi berhasil")
}

// login memverifikasi kredensial user dan mengembalikan role
// Parameter: nUser int - jumlah user saat ini
// Return: string - role user ("admin"/"kasir") atau "" jika gagal
func login(nUser int) string {
	var username, password string

	fmt.Println("\n*** LOGIN ***")
	fmt.Print("Username : ")
	fmt.Scan(&username)
	fmt.Print("Password : ")
	fmt.Scan(&password)

	for i := 0; i < nUser; i++ {
		if dataUser[i].Username == username && dataUser[i].Password == password {
			fmt.Println("Login berhasil sebagai", dataUser[i].Role)
			return dataUser[i].Role
		}
	}

	fmt.Println("Login gagal")
	return ""
}

// tambahPasien menambahkan data pasien baru ke array dataPasien
// Parameter: nPasien *int - pointer ke jumlah pasien saat ini
func tambahPasien(nPasien *int) {
	if *nPasien >= NMAX {
		fmt.Println("Data penuh")
		return
	}

	fmt.Println("\n*** TAMBAH PASIEN ***")
	fmt.Print("ID Pasien : ")
	fmt.Scan(&dataPasien[*nPasien].ID)
	fmt.Print("Nama : ")
	fmt.Scan(&dataPasien[*nPasien].Nama)
	fmt.Print("Alamat : ")
	fmt.Scan(&dataPasien[*nPasien].Alamat)
	fmt.Print("No HP : ")
	fmt.Scan(&dataPasien[*nPasien].NoHP)

	(*nPasien)++
	fmt.Println("Pasien berhasil ditambahkan")
}

// tampilPasien menampilkan seluruh data pasien
// Parameter: nPasien int - jumlah pasien saat ini
func tampilPasien(nPasien int) {
	fmt.Println("\n*** DATA PASIEN ***")

	if nPasien == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i := 0; i < nPasien; i++ {
		fmt.Println("----------------------")
		fmt.Println("ID     :", dataPasien[i].ID)
		fmt.Println("Nama   :", dataPasien[i].Nama)
		fmt.Println("Alamat :", dataPasien[i].Alamat)
		fmt.Println("No HP  :", dataPasien[i].NoHP)
	}
	fmt.Println("----------------------")
}

// sequentialSearchIndexPasien mencari index pasien berdasarkan ID (Sequential Search)
// Parameter: id int - ID yang dicari, nPasien int - jumlah pasien
// Return: int - index jika ditemukan, -1 jika tidak ditemukan
func sequentialSearchIndexPasien(id int, nPasien int) int {
	for i := 0; i < nPasien; i++ {
		if dataPasien[i].ID == id {
			return i
		}
	}
	return -1
}

// sequentialSearchIndexLayanan mencari index layanan berdasarkan ID (Sequential Search)
// Parameter: id int - ID yang dicari, nLayanan int - jumlah layanan
// Return: int - index jika ditemukan, -1 jika tidak ditemukan
func sequentialSearchIndexLayanan(id int, nLayanan int) int {
	for i := 0; i < nLayanan; i++ {
		if dataLayanan[i].ID == id {
			return i
		}
	}
	return -1
}


// ubahPasien mengubah data pasien berdasarkan ID menggunakan sequential search
// Parameter: nPasien int - jumlah pasien saat ini
func ubahPasien(nPasien int) {
	var id int

	fmt.Print("Masukkan ID Pasien : ")
	fmt.Scan(&id)

	idx := sequentialSearchIndexPasien(id, nPasien)
	if idx == -1 {
		fmt.Println("Pasien tidak ditemukan")
		return
	}

	fmt.Print("Nama Baru : ")
	fmt.Scan(&dataPasien[idx].Nama)
	fmt.Print("Alamat Baru : ")
	fmt.Scan(&dataPasien[idx].Alamat)
	fmt.Print("No HP Baru : ")
	fmt.Scan(&dataPasien[idx].NoHP)

	fmt.Println("Data berhasil diubah")
}

// hapusPasien menghapus data pasien berdasarkan ID menggunakan sequential search
// Parameter: nPasien *int - pointer ke jumlah pasien saat ini
func hapusPasien(nPasien *int) {
	var id int

	fmt.Print("Masukkan ID Pasien : ")
	fmt.Scan(&id)

	idx := sequentialSearchIndexPasien(id, *nPasien)
	if idx == -1 {
		fmt.Println("Pasien tidak ditemukan")
		return
	}

	for j := idx; j < *nPasien-1; j++ {
		dataPasien[j] = dataPasien[j+1]
	}
	(*nPasien)--
	fmt.Println("Data berhasil dihapus")
}

// tambahLayanan menambahkan data layanan baru ke array dataLayanan
// Parameter: nLayanan *int - pointer ke jumlah layanan saat ini
func tambahLayanan(nLayanan *int) {
	if *nLayanan >= NMAX {
		fmt.Println("Data penuh")
		return
	}

	fmt.Println("\n*** TAMBAH LAYANAN ***")
	fmt.Print("ID Layanan : ")
	fmt.Scan(&dataLayanan[*nLayanan].ID)
	fmt.Print("Nama Layanan : ")
	fmt.Scan(&dataLayanan[*nLayanan].Nama)
	fmt.Print("Harga : ")
	fmt.Scan(&dataLayanan[*nLayanan].Harga)

	(*nLayanan)++
	fmt.Println("Layanan berhasil ditambahkan")
}

// tampilLayanan menampilkan seluruh data layanan
// Parameter: nLayanan int - jumlah layanan saat ini
func tampilLayanan(nLayanan int) {
	fmt.Println("\n*** DATA LAYANAN ***")

	if nLayanan == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i := 0; i < nLayanan; i++ {
		fmt.Println("----------------------")
		fmt.Println("ID    :", dataLayanan[i].ID)
		fmt.Println("Nama  :", dataLayanan[i].Nama)
		fmt.Println("Harga :", dataLayanan[i].Harga)
	}
	fmt.Println("----------------------")
}

// ubahLayanan mengubah data layanan berdasarkan ID menggunakan sequential search
// Parameter: nLayanan int - jumlah layanan saat ini
func ubahLayanan(nLayanan int) {
	var id int

	fmt.Print("Masukkan ID Layanan : ")
	fmt.Scan(&id)

	idx := sequentialSearchIndexLayanan(id, nLayanan)
	if idx == -1 {
		fmt.Println("Layanan tidak ditemukan")
		return
	}

	fmt.Print("Nama Baru : ")
	fmt.Scan(&dataLayanan[idx].Nama)
	fmt.Print("Harga Baru : ")
	fmt.Scan(&dataLayanan[idx].Harga)

	fmt.Println("Layanan berhasil diubah")
}

// hapusLayanan menghapus data layanan berdasarkan ID menggunakan sequential search
// Parameter: nLayanan *int - pointer ke jumlah layanan saat ini
func hapusLayanan(nLayanan *int) {
	var id int

	fmt.Print("Masukkan ID Layanan : ")
	fmt.Scan(&id)

	idx := sequentialSearchIndexLayanan(id, *nLayanan)
	if idx == -1 {
		fmt.Println("Layanan tidak ditemukan")
		return
	}

	for j := idx; j < *nLayanan-1; j++ {
		dataLayanan[j] = dataLayanan[j+1]
	}
	(*nLayanan)--
	fmt.Println("Layanan berhasil dihapus")
}

// menuAdminPart1 menampilkan menu pengelolaan pasien dan layanan
// Parameter: nPasien *int, nLayanan *int - pointer ke counter masing-masing
func menuAdminPart1(nPasien *int, nLayanan *int) {
	var pilih int

	for {
		fmt.Println("\n*** MENU ADMIN (PART 1) ***")
		fmt.Println("1. Tambah Pasien")
		fmt.Println("2. Tampil Pasien")
		fmt.Println("3. Ubah Pasien")
		fmt.Println("4. Hapus Pasien")
		fmt.Println("5. Tambah Layanan")
		fmt.Println("6. Tampil Layanan")
		fmt.Println("7. Ubah Layanan")
		fmt.Println("8. Hapus Layanan")
		fmt.Println("9. Lanjut Part 2")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahPasien(nPasien)
		case 2:
			tampilPasien(*nPasien)
		case 3:
			ubahPasien(*nPasien)
		case 4:
			hapusPasien(nPasien)
		case 5:
			tambahLayanan(nLayanan)
		case 6:
			tampilLayanan(*nLayanan)
		case 7:
			ubahLayanan(*nLayanan)
		case 8:
			hapusLayanan(nLayanan)
		case 9:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// sequentialSearchNama mencari pasien berdasarkan nama (Sequential Search)
// Parameter: nama string - nama yang dicari, nPasien int - jumlah pasien
func sequentialSearchNama(nama string, nPasien int) {
	for i := 0; i < nPasien; i++ {
		if dataPasien[i].Nama == nama {
			fmt.Println("Pasien ditemukan:")
			fmt.Println("ID     :", dataPasien[i].ID)
			fmt.Println("Nama   :", dataPasien[i].Nama)
			fmt.Println("Alamat :", dataPasien[i].Alamat)
			fmt.Println("No HP  :", dataPasien[i].NoHP)
			return
		}
	}
	fmt.Println("Pasien tidak ditemukan")
}

// sortPasienByID mengurutkan array dataPasien berdasarkan ID secara ascending
// Parameter: nPasien int - jumlah pasien
func sortPasienByID(nPasien int) {
	var temp Pasien

	for i := 0; i < nPasien-1; i++ {
		for j := i + 1; j < nPasien; j++ {
			if dataPasien[i].ID > dataPasien[j].ID {
				temp = dataPasien[i]
				dataPasien[i] = dataPasien[j]
				dataPasien[j] = temp
			}
		}
	}
}

// binarySearchID mencari pasien berdasarkan ID (Binary Search)
// Prasyarat: dataPasien sudah diurutkan berdasarkan ID (panggil sortPasienByID terlebih dahulu)
// Parameter: id int - ID yang dicari, nPasien int - jumlah pasien
func binarySearchID(id int, nPasien int) {
	left := 0
	right := nPasien - 1

	for left <= right {
		mid := (left + right) / 2

		if dataPasien[mid].ID == id {
			fmt.Println("Pasien ditemukan:")
			fmt.Println("ID     :", dataPasien[mid].ID)
			fmt.Println("Nama   :", dataPasien[mid].Nama)
			fmt.Println("Alamat :", dataPasien[mid].Alamat)
			fmt.Println("No HP  :", dataPasien[mid].NoHP)
			return
		} else if dataPasien[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Println("Pasien tidak ditemukan")
}

// tambahTransaksi menambahkan transaksi baru dan otomatis mencatat kunjungan
// Parameter: nTransaksi *int, nKunjungan *int - pointer ke counter masing-masing
func tambahTransaksi(nTransaksi *int, nKunjungan *int) {
	if *nTransaksi >= NMAX {
		fmt.Println("Data penuh")
		return
	}

	fmt.Println("\n*** TAMBAH TRANSAKSI ***")
	fmt.Print("ID Pasien : ")
	fmt.Scan(&dataTransaksi[*nTransaksi].IDPasien)
	fmt.Print("Nama Pasien : ")
	fmt.Scan(&dataTransaksi[*nTransaksi].NamaPasien)
	fmt.Print("Layanan : ")
	fmt.Scan(&dataTransaksi[*nTransaksi].Layanan)
	fmt.Print("Tanggal (angka) : ")
	fmt.Scan(&dataTransaksi[*nTransaksi].Tanggal)
	fmt.Print("Total Biaya : ")
	fmt.Scan(&dataTransaksi[*nTransaksi].TotalBiaya)

	dataKunjungan[*nKunjungan].IDPasien = dataTransaksi[*nTransaksi].IDPasien
	dataKunjungan[*nKunjungan].NamaPasien = dataTransaksi[*nTransaksi].NamaPasien
	dataKunjungan[*nKunjungan].Layanan = dataTransaksi[*nTransaksi].Layanan
	dataKunjungan[*nKunjungan].Tanggal = dataTransaksi[*nTransaksi].Tanggal

	(*nTransaksi)++
	(*nKunjungan)++

	fmt.Println("Transaksi berhasil ditambahkan")
}

// tampilTransaksi menampilkan seluruh data transaksi
// Parameter: nTransaksi int - jumlah transaksi saat ini
func tampilTransaksi(nTransaksi int) {
	fmt.Println("\n*** DATA TRANSAKSI ***")

	if nTransaksi == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i := 0; i < nTransaksi; i++ {
		fmt.Println("----------------------")
		fmt.Println("ID Pasien    :", dataTransaksi[i].IDPasien)
		fmt.Println("Nama Pasien  :", dataTransaksi[i].NamaPasien)
		fmt.Println("Layanan      :", dataTransaksi[i].Layanan)
		fmt.Println("Tanggal      :", dataTransaksi[i].Tanggal)
		fmt.Println("Total Biaya  :", dataTransaksi[i].TotalBiaya)
	}
	fmt.Println("----------------------")
}

// selectionSortTanggal mengurutkan transaksi berdasarkan tanggal (Selection Sort)
// Parameter: nTransaksi int - jumlah transaksi, ascending bool - urutan naik/turun
func selectionSortTanggal(nTransaksi int, ascending bool) {
	var idxEkstrem int
	var temp Transaksi

	for i := 0; i < nTransaksi-1; i++ {
		idxEkstrem = i

		for j := i + 1; j < nTransaksi; j++ {
			if ascending {
				if dataTransaksi[j].Tanggal < dataTransaksi[idxEkstrem].Tanggal {
					idxEkstrem = j
				}
			} else {
				if dataTransaksi[j].Tanggal > dataTransaksi[idxEkstrem].Tanggal {
					idxEkstrem = j
				}
			}
		}

		temp = dataTransaksi[i]
		dataTransaksi[i] = dataTransaksi[idxEkstrem]
		dataTransaksi[idxEkstrem] = temp
	}

	if ascending {
		fmt.Println("Data berhasil diurutkan berdasarkan tanggal (Ascending)")
	} else {
		fmt.Println("Data berhasil diurutkan berdasarkan tanggal (Descending)")
	}
}

// insertionSortBiaya mengurutkan transaksi berdasarkan total biaya (Insertion Sort)
// Parameter: nTransaksi int - jumlah transaksi, ascending bool - urutan naik/turun
func insertionSortBiaya(nTransaksi int, ascending bool) {
	var temp Transaksi
	var j int

	for i := 1; i < nTransaksi; i++ {
		temp = dataTransaksi[i]
		j = i - 1

		if ascending {
			for j >= 0 && dataTransaksi[j].TotalBiaya > temp.TotalBiaya {
				dataTransaksi[j+1] = dataTransaksi[j]
				j--
			}
		} else {
			for j >= 0 && dataTransaksi[j].TotalBiaya < temp.TotalBiaya {
				dataTransaksi[j+1] = dataTransaksi[j]
				j--
			}
		}

		dataTransaksi[j+1] = temp
	}

	if ascending {
		fmt.Println("Data berhasil diurutkan berdasarkan biaya (Ascending)")
	} else {
		fmt.Println("Data berhasil diurutkan berdasarkan biaya (Descending)")
	}
}

// tampilKunjungan menampilkan seluruh riwayat kunjungan pasien
// Parameter: nKunjungan int - jumlah kunjungan saat ini
func tampilKunjungan(nKunjungan int) {
	fmt.Println("\n*** RIWAYAT KUNJUNGAN ***")

	if nKunjungan == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i := 0; i < nKunjungan; i++ {
		fmt.Println("----------------------")
		fmt.Println("ID Pasien   :", dataKunjungan[i].IDPasien)
		fmt.Println("Nama Pasien :", dataKunjungan[i].NamaPasien)
		fmt.Println("Layanan     :", dataKunjungan[i].Layanan)
		fmt.Println("Tanggal     :", dataKunjungan[i].Tanggal)
	}
	fmt.Println("----------------------")
}

// statistik menampilkan statistik kunjungan dan layanan terbanyak
// Parameter: nKunjungan int - jumlah kunjungan saat ini
func statistik(nKunjungan int) {
	var layananTerbanyak string
	var maxCount, count int

	for i := 0; i < nKunjungan; i++ {
		count = 0
		for j := 0; j < nKunjungan; j++ {
			if dataKunjungan[i].Layanan == dataKunjungan[j].Layanan {
				count++
			}
		}

		if count > maxCount {
			maxCount = count
			layananTerbanyak = dataKunjungan[i].Layanan
		}
	}

	fmt.Println("\n*** STATISTIK ***")
	fmt.Println("Jumlah Kunjungan :", nKunjungan)

	if nKunjungan > 0 {
		fmt.Println("Layanan Terbanyak :", layananTerbanyak)
	}
}

// menuAdmin menampilkan dan mengelola menu khusus admin
// Parameter: nPasien *int, nLayanan *int, nKunjungan *int - pointer ke counter
func menuAdmin(nPasien *int, nLayanan *int, nKunjungan *int) {
	var pilih int
	var nama string
	var id int

	menuAdminPart1(nPasien, nLayanan)

	for {
		fmt.Println("\n*** MENU ADMIN (PART 2) ***")
		fmt.Println("1. Search Nama (Sequential Search)")
		fmt.Println("2. Sort Pasien by ID (persiapan Binary Search)")
		fmt.Println("3. Search ID (Binary Search)")
		fmt.Println("4. Statistik")
		fmt.Println("0. Logout")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Print("Masukkan Nama : ")
			fmt.Scan(&nama)
			sequentialSearchNama(nama, *nPasien)
		case 2:
			sortPasienByID(*nPasien)
			fmt.Println("Data pasien berhasil diurutkan berdasarkan ID")
		case 3:
			fmt.Print("Masukkan ID : ")
			fmt.Scan(&id)
			binarySearchID(id, *nPasien)
		case 4:
			statistik(*nKunjungan)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// menuKasir menampilkan dan mengelola menu khusus kasir
// Parameter: nTransaksi *int, nKunjungan *int - pointer ke counter masing-masing
func menuKasir(nTransaksi *int, nKunjungan *int) {
	var pilih int
	var urutPilih int

	for {
		fmt.Println("\n*** MENU KASIR ***")
		fmt.Println("1. Tambah Transaksi")
		fmt.Println("2. Tampil Transaksi")
		fmt.Println("3. Sort Tanggal (Selection Sort)")
		fmt.Println("4. Sort Biaya (Insertion Sort)")
		fmt.Println("5. Tampil Riwayat Kunjungan")
		fmt.Println("0. Logout")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahTransaksi(nTransaksi, nKunjungan)
		case 2:
			tampilTransaksi(*nTransaksi)
		case 3:
			fmt.Println("Urutan :")
			fmt.Println("1. Ascending")
			fmt.Println("2. Descending")
			fmt.Print("Pilih [1-2] : ")
			fmt.Scan(&urutPilih)
			if urutPilih == 1 {
				selectionSortTanggal(*nTransaksi, true)
			} else if urutPilih == 2 {
				selectionSortTanggal(*nTransaksi, false)
			} else {
				fmt.Println("Pilihan tidak valid")
			}
		case 4:
			fmt.Println("Urutan :")
			fmt.Println("1. Ascending")
			fmt.Println("2. Descending")
			fmt.Print("Pilih [1-2] : ")
			fmt.Scan(&urutPilih)
			if urutPilih == 1 {
				insertionSortBiaya(*nTransaksi, true)
			} else if urutPilih == 2 {
				insertionSortBiaya(*nTransaksi, false)
			} else {
				fmt.Println("Pilihan tidak valid")
			}
		case 5:
			tampilKunjungan(*nKunjungan)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func main() {
	var pilih int
	var role string

	// Counter sebagai variabel lokal, bukan global
	var nUser, nPasien, nLayanan, nTransaksi, nKunjungan int

	for {
		banner()
		menuUtama()
		fmt.Print("Pilih menu [1-3] : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			registrasi(&nUser)
		case 2:
			role = login(nUser)
			if role == "admin" {
				menuAdmin(&nPasien, &nLayanan, &nKunjungan)
			} else if role == "kasir" {
				menuKasir(&nTransaksi, &nKunjungan)
			}
		case 3:
			fmt.Println("Terima kasih telah menggunakan SIM-KLIK")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}