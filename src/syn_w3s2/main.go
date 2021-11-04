package main

import (
	"fmt"
	"time"
)

var ItemsChannel = make(chan string)
var CleanedItemsChannel = make(chan string)

func main() {
	items := [10]string{
		"emas",
		"batu",
		"karang",
		"emas",
		"besi",
		"emas",
		"ikan",
		"emas",
		"emas",
		"ikan",
	}

	go menyelam(items)
	go membersihkan()
	go menyimpan()

	time.Sleep(5 * time.Second) // bukan cara best practice untuk nunggu go routine
}

func menyelam(items [10]string) {
	for _, item := range items {
		if item == "emas" {
			fmt.Println("Berhasil menemukan ", item)
			ItemsChannel <- item
		}
	}
}

func membersihkan() {
	for itemKotor := range ItemsChannel {
		fmt.Println("Berhasil membersihkan ", itemKotor)
		CleanedItemsChannel <- itemKotor
	}
}

func menyimpan() {
	for cleanedItem := range CleanedItemsChannel {
		fmt.Println("Berhasil menyimpan ", cleanedItem)
	}
}
