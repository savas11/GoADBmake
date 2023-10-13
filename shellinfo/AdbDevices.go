package shellinfo

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var DeviceID []string

func OpenAdb() {
	cmd := exec.Command("adb", "devices")

	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Komut çalıştırma hatası:", err)
		return
	}

	reading := strings.Split(string(output), "\n")

	if len(reading) < 2 || len(reading[1]) == 0 {
		fmt.Println("Bağlı Cihaz Yoktur!!!")
		os.Exit(0)
	}

	for i := 1; i < len(reading); i++ {
		parts := strings.Fields(reading[i])
		if len(parts) > 1 {
			DeviceID = append(DeviceID, parts[0])
		}
	}
	fmt.Println(len(DeviceID), "cihaz baglidir.")
	fmt.Println(DeviceID, "cihaz baglidir.")
}

func DevicesFor(i int) (int, string) {
	if i < len(DeviceID) {
		return 1, DeviceID[i]
	} else {
		return 0, "nil"
	}
}
