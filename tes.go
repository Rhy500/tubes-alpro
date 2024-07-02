package main

import (
	"fmt"
	"sort"
)

const NMAX = 10

type parawisata struct {
	ID, Tempat, Fasilitas string
	Harga, Jarak          int
}

type tabInt [NMAX]parawisata

func main() {
	var data tabInt
	var nData int

	fmt.Println("Selamat Datang di aplikasi Parawisata")

	bacaData(&data, &nData)
	cetakData(data, nData)

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Tambah data")
		fmt.Println("2. Ubah data")
		fmt.Println("3. Hapus data")
		fmt.Println("4. Urutkan data")
		fmt.Println("5. Keluar")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahData(&data, &nData)
			cetakData(data, nData)
		case 2:
			ubahData(&data, &nData)
			cetakData(data, nData)
		case 3:
			hapusData(&data, &nData)
			cetakData(data, nData)
		case 4:
			urutkanData(&data, &nData)
			cetakData(data, nData)
		case 5:
			fmt.Println("Terima kasih telah menggunakan aplikasi Parawisata.")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func bacaData(A *tabInt, n *int) {
	fmt.Println("Masukkan data pariwisata (ID, Tempat, Fasilitas, Jarak, Harga):")
	for i := 0; i < NMAX; i++ {
		fmt.Scan(&A[i].ID)
		if A[i].ID == "none" {
			break
		}
		fmt.Scan(&A[i].Tempat, &A[i].Fasilitas, &A[i].Jarak, &A[i].Harga)
		*n++
	}
}

func cetakData(A tabInt, n int) {
	fmt.Println("=====================================")
	fmt.Println("             Data Parawisata          ")
	fmt.Println("=====================================")

	for i := 0; i < n; i++ {
		fmt.Printf("  %-5s%-20s%-7s%-5d%-5d\n", A[i].ID, A[i].Tempat, A[i].Fasilitas, A[i].Jarak, A[i].Harga)
	}

	fmt.Println("=====================================")
}

func cetakDataUrut(A tabInt, n int, kriteria string) {
	fmt.Printf("=====================================\n")
	fmt.Printf("         Data urut %s\n", kriteria)
	fmt.Printf("=====================================\n")

	for i := 0; i < n; i++ {
		fmt.Printf("  %-5s%-20s%-7s%-5d%-5d\n", A[i].ID, A[i].Tempat, A[i].Fasilitas, A[i].Jarak, A[i].Harga)
	}

	fmt.Printf("=====================================\n")
}

func tambahData(A *tabInt, n *int) {
	if *n < NMAX {
		fmt.Println("Masukkan data baru (ID, Tempat, Fasilitas, Jarak, Harga):")
		fmt.Scan(&A[*n].ID, &A[*n].Tempat, &A[*n].Fasilitas, &A[*n].Jarak, &A[*n].Harga)
		*n++
		fmt.Println("Data berhasil ditambahkan")
	} else {
		fmt.Println("Array sudah penuh")
	}
}

func ubahData(A *tabInt, n *int) {
	fmt.Println("Masukkan ID data yang ingin diubah:")
	var id string
	fmt.Scan(&id)

	for i := 0; i < *n; i++ {
		if A[i].ID == id {
			fmt.Println("Masukkan data baru (Tempat, Fasilitas, Jarak, Harga):")
			fmt.Scan(&A[i].Tempat, &A[i].Fasilitas, &A[i].Jarak, &A[i].Harga)
			fmt.Println("Data berhasil diubah")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func hapusData(A *tabInt, n *int) {
	fmt.Println("Masukkan ID data yang ingin dihapus:")
	var id string
	fmt.Scan(&id)

	for i := 0; i < *n; i++ {
		if A[i].ID == id {
			for j := i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			*n--
			fmt.Println("Data berhasil dihapus")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func urutkanData(A *tabInt, n *int) {
	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("1. Jarak")
	fmt.Println("2. Biaya")
	fmt.Println("3. Fasilitas")
	var pilihan int
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		sort.Slice(A[:*n], func(i, j int) bool {
			return A[i].Jarak > A[j].Jarak
		})
		cetakDataUrut(*A, *n, "Jarak")
	case 2:
		sort.Slice(A[:*n], func(i, j int) bool {
			return A[i].Harga > A[j].Harga
		})
		cetakDataUrut(*A, *n, "Biaya")
	case 3:
		sort.Slice(A[:*n], func(i, j int) bool {
			return A[i].Fasilitas < A[j].Fasilitas
		})
		cetakDataUrut(*A, *n, "Fasilitas")
	default:
		fmt.Println("Pilihan tidak valid")
		return
	}
}
