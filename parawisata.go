package main

import "fmt"

// Konstanta untuk ukuran maksimum elemen array sebanyak 10
const NMAX = 10

// Tipe bentukan struktur untuk menampung data
type parawisata struct {
	Tujuan, Fasilitas string
	Harga, Jarak      int
}

// tipe alias array of parawisata
type tabInt [NMAX]parawisata

func main() {
	var data tabInt
	var nData int
	var fasilitas, tujuan string
	//var id string

	fmt.Println("Selamat Datang diaplikasi Parawisata")
	// disini tampilan untuk  Pengguna bisa melihat daftar tempat wisata secara terurut berdasarkan kategori-kategori tertentu (misalnya jarak, biaya ataupun fasilitasnya)
	fmt.Println()
	// Baca data
	bacaData(&data, &nData)

	// cetak data awal
	cetakData(data, nData)

	// cari nama snack berdasarkan id
	fmt.Scan(&tujuan)
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

// a. Admin bisa menambahkan, mengubah, dan menghapus data tempat pariwisata pada aplikasi. Data tersebut lengkap dengan failitas dan wahana yang tersedia pada tempat pariwisata tersebut.
func printData(A tabInt, n int) {
	/*
		IS: Array A sebanyak n elemen terdefinisi
		FS: Tercetak di layar data buku dengan format:
			"Data perjalanan :
			 A[i].Tujuan, A[i].Fasilitas, A[i].Jarak, A[i].Harga

	*/
	fmt.Println("Data Perjalanan :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %s %s %d\n", A[i].Tujuan, A[i].Fasilitas, A[i].Jarak, A[i].Harga)
	}
}

func editData(A *tabInt, n int, Fasilitas string, x string) {
	/*
		IS: Array A sebanyak n elemen terdefinisi.
			fasilitas yang akan diedit fasilitas  terdefinisi
			 (x) terdefinisi
		Proses: 1. Lakukan pencarian  berdasarkan tujuan
				   dengan menggunakan fungsi findData.
				2. a. Jika data yang akan diedit berhasil ditemukan, maka
				   lakukan pengubahan  dengan nilai x.
				   Lalu cetak di layar "Data has been edited"
				   b. Jika data yang akan diedit tidak berhasil ditemukan,
				   maka cetak di layar "Data has not been edited".
		FS: Array A sebanyak n elemen berubah isinya, jika data ada
			dalam array (berhasil ditemukan). Jika tidak ada, maka Array tidak
			berubah.
	*/

	idx := findData(*A, n, Fasilitas)
	if idx != -1 {
		A[idx].Fasilitas = x
		fmt.Println("Data has been edited")
	} else {
		fmt.Println("Data has not been edited")
	}
}

func hapusData(A *tabInt, n *int, Fasilitas string) {
	idx := findData(*A, *n, Fasilitas)
	if idx != -1 {
		for i := idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n -= 1
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func findData(A tabInt, n int, fasilitas string) int {
	/* Diberikan array A sebanyak n elemen dan fasilitas dari tujuan, fungsi
	   akan mengembalikan fasilitas itu jika ditemukan, dan mengembalikan
	   -1 jika tidak ditemukan. Catatan: Gunakan algoritma sequential search */
	var i, idx int

	idx = -1
	i = 0

	for i < n && idx == -1 {
		if A[i].Fasilitas == fasilitas {
			idx = i
		}
		i++
	}
	return idx
}

//b. Pengguna bisa melihat daftar tempat wisata secara terurut berdasarkan kategori-kategori tertentu (misalnya jarak, biaya ataupun fasilitasnya)
func bacaData(A *tabInt, n *int) {
	/*
		IS: Array A terdefinisi sembarang
		Proses : Selama array belum penuh dan yang dibaca bukan "none" "none" 0 0,
				 maka data dimasukkan ke dalam array A
		FS: Array A terisi sebanyak n elemen
	*/
	for i := 0; i < NMAX; i++ {
		fmt.Scan(&A[i].Tujuan)
		if A[i].Tujuan == "none" {
			break
		}
		fmt.Scan(&A[i].Fasilitas, &A[i].Harga, &A[i].Jarak)
		*n++
	}
}

func cetakData(A tabInt, n int) {
	/*
		IS: Terdefinisi Array A  sebanyak n elemen
		FS: Tercetak di layar seluruh atribut A dengan format seperti contoh
	*/
	fmt.Println("=====================================")
	fmt.Println("             Data Parawisata         ")
	fmt.Println("=====================================")

	for i := 0; i < n; i++ {
		fmt.Printf("  %-5s%-20s%-7d%-5d\n", A[i].Tujuan, A[i].Fasilitas, A[i].Harga, A[i].Jarak)
	}

	fmt.Println("=====================================")
}

/*func namaSnackSeq(A tabInt, n int, x string) string {
/* Diberikan array A sebanyak n elemen terurut acak, fungsi mengembalikan
   nama snack dengan id (x) jika x terdapat pada array A. Jika tidak
   terdapat pada array A, maka kembalikan string "None" */
/*for i := 0; i < n; i++ {
		if A[i].Tujuan == x {
			return A[i].Fasilitas
		}
	}
	return "None"
}
*/
func omset(A tabInt, n int) int {
	/*
		Diberikan array A sebanyak n elemen, fungsi mengembalikan omset
		yang merupakan jumlah total dari perkalian antara stok dengan harga
	*/
	total := 0
	for i := 0; i < n; i++ {
		total += A[i].Harga * A[i].Jarak
	}
	return total
}

func urutFasilitas(A *tabInt, n int) {
	/*
		IS: Terdefinisi array A sebanyak n elemen dengan terurut sembarang
		FS: Urutkan array A secara menaik (ascending) berdasarkan ID dengan
			algoritma selection sort
	*/
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if (*A)[j].Tujuan < (*A)[minIndex].Tujuan {
				minIndex = j
			}
		}
		(*A)[i], (*A)[minIndex] = (*A)[minIndex], (*A)[i]
	}
}

func urutHarga(A *tabInt, n int) {
	/*
		IS: Terdefinisi array A sebanyak n elemen dengan terurut sembarang
		FS: Urutkan array A secara menurun (descending) berdasarkan harga dengan
		    algoritma insertion sort
	*/
	for i := 1; i < n; i++ {
		key := (*A)[i]
		j := i - 1
		for j >= 0 && (*A)[j].Harga < key.Harga {
			(*A)[j+1] = (*A)[j]
			j = j - 1
		}
		(*A)[j+1] = key
	}
}

func urutJarak(A *tabInt, n int) {
	/*
		IS: Terdefinisi array A sebanyak n elemen dengan terurut sembarang
		FS: Urutkan array A secara menurun (descending) berdasarkan harga dengan
		    algoritma insertion sort
	*/
	for i := 1; i < n; i++ {
		key := (*A)[i]
		j := i - 1
		for j >= 0 && (*A)[j].Harga < key.Harga {
			(*A)[j+1] = (*A)[j]
			j = j - 1
		}
		(*A)[j+1] = key
	}
}

func namaHargaBin(A tabInt, n int, x int) string {
	/* Diberikan array A sebanyak n elemen terurut menurun berdasarkan harga,
	   fungsi mengembalikan fasilitas dengan harga (x) jika x terdapat pada
	   array A. Jika tidak terdapat pada array A, maka kembalikan string "None"
	*/
	low, high := 0, n-1
	for low <= high {
		mid := low + (high-low)/2
		if A[mid].Harga == x {
			return A[mid].Fasilitas
		} else if A[mid].Harga < x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return "None"
}
