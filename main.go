package main

import (
	"GoADBmake/choise"
	"GoADBmake/shellinfo"
	"fmt"
)

func main() {
	shellinfo.OpenAdb()
	var text string
	fmt.Printf("Lutfen Kullanici adini bosluk olmadan giriniz : ")
	fmt.Scanf("%s", &text)
	fmt.Println(" ")
	choise.MainChoise(text)
}
