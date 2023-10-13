package platform

import (
	"GoADBmake/devicedisplay"
	"GoADBmake/shellinfo"
	"fmt"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

func InstaLiketoLike(text string) {
	var wg sync.WaitGroup

	for i := 0; i < len(shellinfo.DeviceID); i++ {
		durum, id := shellinfo.DevicesFor(i)
		if durum == 1 {
			wg.Add(1)
			go func(id string, text string, i int) {
				defer wg.Done()
				devicedisplay.UnLockTapSwipe(id)
				if OpenInsta(id) == 1 {
					SearchButton(id)       //Anasayfa alt arama butonu
					SearchToolbarTouch(id) //Ust arama cubuguna dokunma
					UserText(id, text)     //kullanici adi yazma
					time.Sleep(2 * time.Second)
					SearchUserButton(id) //kullanici adini genel arama yapma
					NearSwipe(id)        //hesaplar bolumune gecis icin yana kaydirir
					time.Sleep(2 * time.Second)
					UserTouchButton(id) //Ilk kullanici icin acik olsun
					//UserTouchButton2(id) //Ikinci kullanici icin acik olsun
					PhotoTouchButton(id) //fotoya tiklama
					time.Sleep(2 * time.Second)
					LikeDoubleTouch(id) //likela
					fmt.Println("ID'si", id, "olan", i, "cihazin islemi bitmistir.")
					time.Sleep(5 * time.Second)
					CloseInsta(id)            //instayi kapat
					devicedisplay.LockTap(id) //ekrani kitle
				}
			}(id, text, i)
		}
	}

	wg.Wait()
	fmt.Println("Tum Cihazlara LiketoLike ozelligi saglanmistir")
}

func OpenInsta(id string) int {
	cmd := exec.Command("adb", "-s", id, "shell", "am", "start", "-n", "com.instagram.android/com.instagram.android.activity.MainTabActivity")

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Instagram acilma hatasi: %v\n", err)
		return 0
	}
	time.Sleep(5 * time.Second)
	return 1
}

func SearchButton(id string) {
	x, y := devicedisplay.DisplaySize(id)
	a := int(float64(x) / float64(3.55))
	b := int(float64(y) - float64(y)/float64(16.40))
	cmd := exec.Command("adb", "-s", id, "shell", "input", "tap", strconv.Itoa(a), strconv.Itoa(b))
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Kilit Tusu Hatasi: %v\n", err)
		return
	}
	time.Sleep(3 * time.Second)
}

func SearchToolbarTouch(id string) {
	x, y := devicedisplay.DisplaySize(id)
	a := int(float64(x) / float64(3.55))
	b := int(float64(y) / float64(12))
	cmd := exec.Command("adb", "-s", id, "shell", "input", "tap", strconv.Itoa(a), strconv.Itoa(b))
	err := cmd.Run()
	if err != nil {
		fmt.Printf("SearchToolbarTouch: %v\n", err)
		return
	}
}

func UserText(id string, text string) {
	cmd := exec.Command("adb", "-s", id, "shell", "input", "text", text)

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Klavye komut hatasi %v\n", err)
		return
	}
	cmd2 := exec.Command("adb", "-s", id, "shell", "input", "keyevent", "4")

	err2 := cmd2.Run()
	if err != nil {
		fmt.Printf("Klavye geri komut hatasi %v\n", err2)
		return
	}
}

func CloseInsta(id string) int {
	cmd := exec.Command("adb", "-s", id, "shell", "am", "force-stop", "com.instagram.android")

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Instagram kapanma hatasi: %v\n", err)
		return 0
	}
	time.Sleep(5 * time.Second)
	return 1
}

func UserTouchButton(id string) {
	x, y := devicedisplay.DisplaySize(id)
	a := int(float64(x) / float64(2.5)) //100-900
	b := int(float64(y) / float64(5.2)) //400-520
	cmd := exec.Command("adb", "-s", id, "shell", "input", "tap", strconv.Itoa(a), strconv.Itoa(b))
	err := cmd.Run()
	if err != nil {
		fmt.Printf("UserTouchButton Hatasi: %v\n", err)
		return
	}
	time.Sleep(1 * time.Second)
}

func UserTouchButton2(id string) {
	x, y := devicedisplay.DisplaySize(id)
	a := int(float64(x) / float64(2.5)) //100-900
	b := int(float64(y) / float64(3.5)) //634-726
	cmd := exec.Command("adb", "-s", id, "shell", "input", "tap", strconv.Itoa(a), strconv.Itoa(b))
	err := cmd.Run()
	if err != nil {
		fmt.Printf("UserTouchButton Hatasi: %v\n", err)
		return
	}
	time.Sleep(1 * time.Second)
}

func PhotoTouchButton(id string) {
	x, y := devicedisplay.DisplaySize(id)
	a := int(float64(x) / float64(7.2))  //50-300
	b := int(float64(y) / float64(2.19)) //990-1200
	cmd := exec.Command("adb", "-s", id, "shell", "input", "tap", strconv.Itoa(a), strconv.Itoa(b))
	err := cmd.Run()
	if err != nil {
		fmt.Printf("PhotoTouchButton hatasi: %v\n", err)
		return
	}
	time.Sleep(1 * time.Second)
}

func LikeDoubleTouch(id string) {
	cmd := exec.Command("adb", "-s", id, "shell", "input", "tap", "540", "1200")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Like Double Touch hatasi: %v\n", err)
		return
	}
	cmd = exec.Command("adb", "-s", id, "shell", "input", "tap", "550", "1300")
	err2 := cmd.Run()
	if err != nil {
		fmt.Printf("Like Double Touch hatasi: %v\n", err2)
		return
	}

	time.Sleep(2 * time.Second)

	cmd = exec.Command("adb", "-s", id, "shell", "input", "tap", "560", "1350")
	err3 := cmd.Run()
	if err != nil {
		fmt.Printf("Like Double Touch hatasi: %v\n", err3)
		return
	}
	time.Sleep(1 * time.Second)
	cmd = exec.Command("adb", "-s", id, "shell", "input", "tap", "530", "1225")
	err4 := cmd.Run()
	if err != nil {
		fmt.Printf("Like Double Touch hatasi: %v\n", err4)
		return
	}
	cmd = exec.Command("adb", "-s", id, "shell", "input", "tap", "555", "1205")
	err5 := cmd.Run()
	if err != nil {
		fmt.Printf("Like Double Touch hatasi: %v\n", err5)
		return
	}
}

func SearchUserButton(id string) {
	x, y := devicedisplay.DisplaySize(id)
	a := int(float64(x) / float64(2)) //80-935
	b := int(float64(y) / float64(6)) //200-420
	cmd := exec.Command("adb", "-s", id, "shell", "input", "tap", strconv.Itoa(a), strconv.Itoa(b))
	err := cmd.Run()
	if err != nil {
		fmt.Printf("UserTouchButton Hatasi: %v\n", err)
		return
	}
	time.Sleep(1 * time.Second)
}

func NearSwipe(deviceID string) {
	cmd := exec.Command("adb", "-s", deviceID, "shell", "input ", "swipe", "900", "1200", "300", "1200", "100")
	err2 := cmd.Run()
	if err2 != nil {
		fmt.Printf("Ekran Kaydirma Hareketi: %v\n", err2)
		return
	}
}
