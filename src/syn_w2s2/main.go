package main

import "fmt"

func StartSection(value string) {
	fmt.Printf("\n===================%s===================\n", value)
}

// fucntion biasa
// func PenjumlahanAngka(angkaPertama, angkaKedua, angkaKetiga int) {

// }

// ================================================================
// variadic function => variabel argumen
func PenjumlahanAngka(angka ...int) int {
	total := 0
	for _, value := range angka { // angka berubah menjadi slice tipe integer
		total += value
	}

	return total
}

// ================================================================
// var argumen harus dibelakang
func belanjaBulanan(barang string, harga int, diskon ...int) {
	if len(diskon) > 0 {
		harga -= diskon[0]
	}

	fmt.Printf("Saya membeli barang %s dengan total belanja adalah %d\n", barang, harga)
}

// ================================================================

// anonymous function
// sistem auth sederhana
// ide = buat sebuah function untuk cek value berupa nama
// hasil = user akan di block, user akan dapat akses

type Blacklist func(string) bool

func loginUser(nama string, blacklist Blacklist) {
	if blacklist(nama) {
		fmt.Println("Mohon maaf anda tidak dapat akses", nama)
	} else {
		fmt.Println("Selamat datang", nama)
	}
}

// ===========================Defer=====================================
func logging() {
	fmt.Println("Selesai di jalankan")
}

func pembagian(value int) {
	defer logging()
	fmt.Println("Fungsi Pembagian")
	hasil := 10 / value
	fmt.Println("Hasilnya adalah", hasil)
}

// ===========================Struct=====================================
type Mahasiswa struct {
	Nama, ProgramStudi string
	Semester, Umur     int
	StatusAktif        bool
}

// ===========================Struct Method=====================================
func (mahasiswa Mahasiswa) grandOpening(namaDosen string) {
	fmt.Println("Hello perkenalkan saya dosen", namaDosen, "selamat datang", mahasiswa.Nama)
}

// ===========================Anonymous Struct=====================================
func anonymousStruct() {
	MataKuliah := struct {
		// Tempat deklarasi field struct
		Subject string
		Credit  int
	}{
		// Tempat untuk isi data struct
		Subject: "Alstrudat",
		Credit:  3,
	}

	if MataKuliah.Credit > 2 {
		fmt.Println("Mata Kuliah", MataKuliah.Subject, "Harus belajar lebih giat")
	} else {
		fmt.Println("Mata Kuliah", MataKuliah.Subject, "bisa sedikit santai")
	}
}

func main() {
	StartSection("Variabel Argumen")
	sliceParameter := []int{3, 4, 5, 6, 7, 8}
	r := PenjumlahanAngka(sliceParameter...)
	fmt.Println(r)

	StartSection("Variabel Argumen")
	belanjaBulanan("susu", 45000)
	belanjaBulanan("detergen", 15000, 2000)

	StartSection("Anonymous Function")
	blokir := func(nama string) bool {
		return nama == "admin"
	}

	loginUser("admin", blokir)

	StartSection("Anonymous Function")
	loginUser("kornelius", func(nama string) bool {
		return nama == "kornelius"
	})

	StartSection("Defer")
	// defer fmt.Println("World")
	// defer fmt.Println("Hello")
	// Output: Hello World
	fmt.Println("Function dijalankan")
	pembagian(2)

	StartSection("Struct")
	student := Mahasiswa{
		Nama:         "Kornelius",
		ProgramStudi: "Teknologi Informasi",
		Semester:     5,
		Umur:         20,
		StatusAktif:  true,
	}
	fmt.Println(student)

	StartSection("Struct Method")
	studentNew := Mahasiswa{Nama: "Kornelius"}
	studentNew.grandOpening("Budi")

	StartSection("Anonymous Struct")
	anonymousStruct()
}
