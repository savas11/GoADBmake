package choise

import (
	"GoADBmake/platform"
	"fmt"
	"os"
)

func MainChoise(text string) {

	fmt.Println("******************************************")
	fmt.Println("*--Hangi Uygulama kullanmak istersiniz?--*")
	fmt.Println("*--1. Instagram -------------------------*")
	fmt.Println("*--2. Facebook --------------------------*")
	fmt.Println("*--3. Youtube ---------------------------*")
	fmt.Println("*----------------------------------------*")
	fmt.Println("*--4. Exit ------------------------------*")
	fmt.Println("******************************************")
	fmt.Printf("Lutfen Sayi Giriniz : ")
	var sayi int
	fmt.Scanf("%d", &sayi)
	fmt.Println(" ")
	switch sayi {
	case 1:
		InstaOrientation(text)
	case 2:
		fmt.Println("Malesef Suan Aktif Degil")
	case 3:
		fmt.Println("Yapim asamasinda")
	case 4:
		os.Exit(0)
	default:
		fmt.Println("Geçersiz seçim.")
	}
}

func InstaOrientation(text string) {
	fmt.Println("******************************************")
	fmt.Println("*--Instagram ozelliklerimiz--------------*")
	fmt.Println("*--1. Like to Like ----------------------*")
	fmt.Println("*--2. Follow ME -------------------------*")
	fmt.Println("*----------------------------------------*")
	fmt.Println("*--3. Exit ------------------------------*")
	fmt.Println("******************************************")
	fmt.Printf("Lutfen Sayi giriniz : ")
	var sayi int
	fmt.Scan(&sayi)
	fmt.Println(" ")
	switch sayi {
	case 1:
		fmt.Println("LiketoLike Isleminiz Baslatiliyor..")
		platform.InstaLiketoLike(text)
		platform.ChromeInstaLiketoLike(text)
	case 2:
		fmt.Println("Daha Aktiflesmedi")
	case 3:
		os.Exit(1)
	default:
		fmt.Println("Geçersiz seçim.")
	}

}
