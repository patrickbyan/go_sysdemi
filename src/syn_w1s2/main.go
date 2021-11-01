package main

import "fmt"

func main() {
	// Dlm function
	var nama string = "alex"
	umur := "12"
	fmt.Printf("Hai nama saya adalah %s sekarang saya berumur %s", nama, umur)

	var namaDepan string = "Agus"
	namaTengah := "Sth"
	var namaBelakang string = "Mulyono"
	fmt.Printf("Hai teman-teman nama saya %s%s%s!\n", namaDepan, namaBelakang, namaTengah)

	// Multi Variable
	var namaa, jenis_kelamin, alamat string
	namaa, jenis_kelamin, alamat = "Agus", "Laki-laki", "Jln Mawar"
	fmt.Println(namaa, jenis_kelamin, alamat, umur)

	name, gender, age := "Agus", "Laki-laki", 20
	fmt.Println(name, gender, age)

	// ================ if else ================
	// Buatlah sebuah program yang akan menampilkan teks 'kamu lulus ujian' ketika nilai
	// diatas dari 70 sedangkan nilai dibawah 70 gagal

	nilaiUjian := 80

	if nilaiUjian > 70 {
		fmt.Println("Selamat kamu lulus ujian")
	} else {
		fmt.Println("Kamu gagal")
	}

	// ================ mapping ================
	// tampilkanlah 10 baris kata "saya belajar golang"
	for i := 0; i < 10; i++ {
		fmt.Println("Saya belajar golang")
	}

	// ================ Array ================
	var examScore [7]int8
	// var examScore [...]int8 // pjg array tidak dideclare sebelumnya

	examScore[0] = 60
	examScore[1] = 70
	examScore[2] = 60
	examScore[3] = 70
	examScore[4] = 60

	// kalau lebih dr index ke 7, udah diluar dari panjang array.
	// kalau i gak digunain, bisa i diganti _
	for _, nilai := range examScore { // range sama dengan len ( panjang ), untuk kapasitas itu cap
		fmt.Println(nilai)
	}

	// ================ slice ================
	// namaMahasiswa := make([]string, 3, 5)
	var hurufVokal = [5]string{"a", "i", "u", "e", "o"}

	var sliceHuruf = hurufVokal[1:3] // hasil: [i u]
	fmt.Println(sliceHuruf)
	fmt.Println(cap(sliceHuruf))
	fmt.Println(len(sliceHuruf))
}
