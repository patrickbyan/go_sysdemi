package main

import "fmt"

func variables() {
	var namaTeman1, namaTeman2, namaTeman3, namaTeman4, namaTeman5 string
	namaTeman1, namaTeman2, namaTeman3, namaTeman4, namaTeman5 = "patrick", "byan", "katamsi", "teman4", "teman5"
	fmt.Printf("Manifest Typing %s %s %s %s %s \n", namaTeman1, namaTeman2, namaTeman3, namaTeman4, namaTeman5)
	teman1, teman2, teman3, teman4, teman5 := "patrick", "byan", "katamsi", "pat", "rick"
	fmt.Printf("Inference %s %s %s %s % s\n", teman1, teman2, teman3, teman4, teman5)
}

var nama string

func variables_fixProblemOne() {
	nama = "rafael"
	jenisKelamin, tinggiBadan, beratBadan := "Laki-laki", 173, 80
	fmt.Printf("Hai nama saya %s\n"+"jenis kelamin saya %s\n"+"tinggi badan saya adalah %d\n"+"dengan berat badan adalah %d \n", nama, jenisKelamin, tinggiBadan, beratBadan)
}

func variables_fixProblemTwo() {
	var nama string = "alex"
	umur := 12
	fmt.Printf("Hai nama saya adalah %s sekarang saya berumur %d \n", nama, umur)
}

func constants() {
	const inputMeter float32 = 20
	const feetToMeter float32 = 0.3048
	const convertMeterToFeet = inputMeter * feetToMeter
	fmt.Printf("%v meter = %v ft \n", inputMeter, convertMeterToFeet)
}

func dataTypes() {
	exp1 := true && false
	exp2 := false && true
	exp3 := false && false
	allExp := exp1 || exp2 || !exp3
	fmt.Println(allExp)
}

func main() {
	variables()
	variables_fixProblemOne()
	variables_fixProblemTwo()
	constants()
	dataTypes()
}
