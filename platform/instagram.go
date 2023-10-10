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
					SearchButton(id)
					SearchToolbarTouch(id)
					UserText(id, text)
					fmt.Println("ID'si", id, "olan", i, "cihazin islemi bitmistir.")
					//CloseInsta(id)
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
		fmt.Printf("Kilit Tusu Hatasi: %v\n", err)
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
