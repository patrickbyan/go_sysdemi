package main

import "fmt"

// func sayHai() {
// 	fmt.Println("Hello, saya sedang belajar golang")
// }

func ToeflMahasiswa() {
	nilaiToefl := [6]int{560, 450, 400, 500, 450}
	sumToefl := 0
	for _, toefl := range nilaiToefl {
		sumToefl += toefl
		//sumToefl = sumToefl + toefl
	}
	avgToefl := sumToefl / len(nilaiToefl)
	fmt.Println("Total jumlah nilai toefl mahasiswa adalah ", sumToefl)
	fmt.Println("Rata-rata nilai toefl mahasiwa adalah ", avgToefl)
}

func Perulangan() {
	nilaiToeflA := [6]int{560, 450, 400, 500, 450}
	for _, data := range nilaiToeflA {
		fmt.Println(data)
	}
}

func sayHello(namaDepan, namaBelakang string) {
	fmt.Println("Hello perkenalkan nama saya adalah ", namaDepan, namaBelakang)
}

func hukumanUjian(nilaiUjian int) {
	if nilaiUjian < 70 {
		for i := 0; i < 10; i++ {
			fmt.Println("saya akan belajar lebih giat lagi")
		}
	} else {
		fmt.Println("selamat kamu lulu ujian")
	}
}

func avgExam(ipa, ips, ppkn, sbk float64) float64 {
	return (ipa + ips + ppkn + sbk) / 4
}

func keuntunganPenjualan(domesticSale, domesticCogs float64) (float64, float64) {
	keuntungan := domesticSale - domesticCogs
	presentaseKeuntungan := (keuntungan / domesticCogs) * 100
	return keuntungan, presentaseKeuntungan
}

func main() {
	// for i := 0; i < 10; i++ {
	// 	sayHai()
	// }

	// Perulangan()
	//ToeflMahasiswa()

	//var namaDepan = "Kornelius"
	//var _ = "sipayung"
	//
	//fmt.Println(namaDepan)

	sayHello("kornelius", "sipayung")

	hukumanUjian(50)
	hukumanUjian(75)

	var (
		ipa  = 60.5
		ips  = 66.6
		ppkn = 70.7
		sbk  = 90.60
	)

	result := avgExam(ipa, ips, ppkn, sbk)
	fmt.Printf("Nilai rata-rata dari kelas X MIA 1 adalah %.2f", result)

	var (
		hargaJual = 125000.50
		hargaBeli = 110000.40
	)
	keuntungan, persentaseKeuntungan := keuntunganPenjualan(hargaJual, hargaBeli)
	fmt.Printf("Keuntungan nya adalah Rp. %.2f dengan persentase Keuntungan %.2f percent", keuntungan, persentaseKeuntungan)
}
