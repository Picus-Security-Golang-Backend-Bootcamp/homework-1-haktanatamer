package main

import (
	"bufio"
	"fmt"
	"os"

	"package.local/hktn/helper"
)

var books []string
var choice int
var reset string
var searchWord string
var counter int

func main() {

	if len(books) == 0 {
		initializeBook()
	}

	fmt.Println("Kitaplık uygulamasında kullanabileceğiniz komutlar :")
	fmt.Println(" search => arama işlemi için 1")
	fmt.Println(" list => listeleme işlemi için 2")
	fmt.Println(" exit => uygulamadan çıkmak için 3")
	fmt.Println("Tuşlarına basınız")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		search()
	case 2:
		listAll()
	case 3:
		terminate()
	default:
		helper.AddSeparator()
		fmt.Println("Yanlış seçim. Tekrar deneyiniz")
		helper.AddSeparator()
		main()
	}
}

func listAll() {
	counter = 0
	fmt.Println("Kitaplar Getiriliyor..")
	helper.AddSeparator()
	for _, line := range books {
		counter++
		helper.ConsoleBook(counter, line)
	}
	helper.AddSeparator()
	fmt.Printf("%d Tane Kitaplar Getirildi..\n", counter)
	helper.AddSeparator()
	conclude()
}

func search() {
	fmt.Println("Aranacak Kelimeyi Giriniz : ")
	fmt.Scan(&searchWord)
	helper.AddSeparator()
	fmt.Printf("%s Kelimesi Aranıyor..\n", searchWord)
	helper.AddSeparator()
	searchByWord(searchWord)
	helper.AddSeparator()
	fmt.Printf("%d Tane Kitap Bulundu..\n", counter)
	fmt.Println("Arama Tamamlandı..")
	helper.AddSeparator()
	conclude()
}

func terminate() {
	fmt.Println("Uygulama Sonlandırılıyor..")
	os.Exit(3)
}

func conclude() {
	fmt.Println("Yeni Bir Arama Yapmak İçin N \nUygulamyı Sonlandırmak İçin Herhangi Bir Tuşa Basınız")

	fmt.Scan(&reset)

	switch reset {
	case "n":
		main()
		break
	case "N":
		main()
		break
	default:
		terminate()
	}
}

func searchByWord(word string) {
	counter = 0
	for _, line := range books {
		if helper.StringContains(line, word) {
			counter++
			helper.ConsoleBook(counter, line)
		}
	}
}

func initializeBook() {
	file, errFile := os.Open("books.txt")
	if errFile != nil {
		helper.AddLog("books.txt okunurken hata alındı.", errFile)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		books = append(books, scanner.Text())
	}
}
