package devicedisplay

import (
	"fmt"
	"os/exec"
	"time"
)

func UnLockTapSwipe(deviceID string) {
	cmd := exec.Command("adb", "-s", deviceID, "shell", "input", "keyevent", "224")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Kilit Tusu Hatasi: %v\n", err)
		return
	}
	time.Sleep(2 * time.Second)
	cmd = exec.Command("adb", "-s", deviceID, "shell", "input ", "swipe", "500", "1800", "500", "800", "100")
	err2 := cmd.Run()
	if err2 != nil {
		fmt.Printf("Ekran Kaydirma Hareketi: %v\n", err2)
		return
	}
}
func LockTap(deviceID string) {
	cmd := exec.Command("adb", "-s", deviceID, "shell", "input", "keyevent", "26")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Kilit Tusu Hatasi: %v\n", err)
		return
	}
}
