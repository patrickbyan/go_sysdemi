package main // indikasi file pertama kali execute

import "fmt"

func main() {
	var namaDepan string = "Agus"

	// Dlm function
	namaTengah := "Sth"

	var namaBelakang string
	namaBelakang = "Mulyono"

	// Multi Variable
	var nama, jenis_kelamin, alamat string
	nama, jenis_kelamin, alamat = "Agus", "Laki-laki", "Jln Mawar"

	nama, jenis_kelamin, umur := "Agus", "Laki-laki", 20

	fmt.Printf("Hai teman-teman nama saya %s %s %s!\n", namaDepan, namaBelakang, namaTengah)

	fmt.Println(nama, jenis_kelamin, alamat, umur)
}
