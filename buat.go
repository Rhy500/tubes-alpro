package main

import "fmt"

// Konstanta untuk ukuran maksimum elemen array sebanyak 10
const NMAX = 10

// Tipe bentukan struktur untuk menampung data
type parawisata struct {
	ID, Tempat, Fasilitas string
	Harga, Jarak      int
}

// tipe alias array of parawisata
type tabInt [NMAX]parawisata

func main() {
	var data tabInt
	var nData int
	var fasilitas string
	var harga, jarak int
	//var id string

	fmt.Println("Selamat Datang diaplikasi Parawisata")
	// disini tampilan untuk  Pengguna bisa melihat daftar tempat wisata secara terurut berdasarkan kategori-kategori tertentu (misalnya jarak, biaya ataupun fasilitasnya)
	fmt.Println()
	// Baca data
	bacaData(&data, &nData)

	// cetak data awal
	cetakData(data, nData)

	// cari nama fasilitas berdasarkan id
	fmt.Scan(&id)
	fmt.Println("parawisata:", namaHargaBin(data, nData, x))

	// Mengurutkan array berdasarkan ID secara menaik
	urutFasilitas(&data, nData)

	// cetak data setelah diurutkan
	cetakData(data, nData)

	// Mengurutkan array berdasarkan harga secara menurun
	urutHarga(&data, nData)

	// cetak data setelah diurutkan
	cetakData(data, nData)

}
func bacaData(A *tabInt, n *int) {
	/*
		IS: Array A terdefinisi sembarang
		Proses : Selama array belum penuh dan yang dibaca bukan "none" "none" 0 0,
				 maka data dimasukkan ke dalam array A
		FS: Array A terisi sebanyak n elemen
	*/
	for i := 0; i < NMAX; i++ {
		fmt.Scan(&A[i].ID)
		if A[i].ID == "none" {
			break
		}
		fmt.Scan(&A[i].Tempat, &A[i].Fasilitas, &A[i].Jarak)
		*n++
	}
}

func cetakData(A tabint, n int) {
	/*
		IS: Terdefinisi Array A  sebanyak n elemen
		FS: Tercetak di layar seluruh atribut A dengan format seperti contoh
	*/
	fmt.Println("=====================================")
	fmt.Println("             Data         ")
	fmt.Println("=====================================")

	for i := 0; i < n; i++ {
		fmt.Printf("  %-5s%-20s%-7d%-5d\n", A[i].ID, A[i].Tempat, A[i].Fasilitas, A[i].Jarak)
	}

	fmt.Println("=====================================")
}
func editData(A *tabInt, n int, ID string, x string) {
	
	idx := findData(*A, n, ID)
	if idx != -1 {
		A[idx].Tempat = x
		fmt.Println("Data has been edited")
	} else {
		fmt.Println("Data has not been edited")
	}
}

func findData(A tabInt, n int,ID string) int {
	/* Diberikan array A sebanyak n elemen dan book Id (bId) dari buku, fungsi
	   akan mengembalikan indeks buku itu jika ditemukan, dan mengembalikan
	   -1 jika tidak ditemukan. Catatan: Gunakan algoritma sequential search */
	var i, idx int

	idx = -1
	i = 0

	for i < n && idx == -1 {
		if A[i].ID == ID {
			idx = i
		}
		i++
	}
	return idx
}