package main

import (
	"GoADBmake/platform"
	"GoADBmake/shellinfo"
)

func main() {
	shellinfo.OpenAdb()
	var text string
	// fmt.Printf("Lutfen Kullanici adini bosluk olmadan giriniz : ")
	// fmt.Scanf("%s", &text)
	// fmt.Println(" ")
	// choise.MainChoise(text)

	//hizli deneme

	text = "eskisehirmekanvlog"
	platform.InstaLiketoLike(text)

}
