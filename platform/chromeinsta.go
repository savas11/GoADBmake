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

func ChromeInstaLiketoLike(text string) {
	var wg sync.WaitGroup

	for i := 0; i < len(shellinfo.DeviceID); i++ {
		durum, id := shellinfo.DevicesFor(i)
		if durum == 1 {
			wg.Add(1)
			go func(id string, text string, i int) {
				defer wg.Done()

				//devicedisplay.UnLockTapSwipe(id)
				if durum == 1 {
					SearchChromeToolbarTouch(id)
					//UserText(id, text)
					fmt.Printf("ID'si %s olan %d. cihazın işlemi tamamlandı.\n", id, i)
				}
			}(id, text, i)
		}
	}

	wg.Wait()
	fmt.Println("Tum Cihazlara LiketoLike ozelligi saglanmistir")
}

func OpenChrome(id string) int {
	cmd := exec.Command("adb", "-s", id, "shell", "am", "start", "-n", "com.android.chrome/com.google.android.apps.chrome.Main")

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Chrome acilma hatasi: %v\n", err)
		return 0
	}
	time.Sleep(5 * time.Second)
	return 1
}

func SearchChromeToolbarTouch(id string) {
	x, y := devicedisplay.DisplaySize(id)
	a := int(float64(x) / float64(2.60)) //670-730
	b := int(float64(y) / float64(4.8))  //120-730
	cmd := exec.Command("adb", "-s", id, "shell", "input", "tap", strconv.Itoa(a), strconv.Itoa(b))
	err := cmd.Run()
	if err != nil {
		fmt.Printf("SearchChromeToolbarTouch: %v\n", err)
		return
	}
}

func UserWebText(id string, text string) {
	website := "instagram.com/" + text
	cmd := exec.Command("adb", "-s", id, "shell", "input", "text", website)

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
