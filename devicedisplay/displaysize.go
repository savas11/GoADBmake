package devicedisplay

import (
	"fmt"
	"os/exec"
	"strconv"
)

func DisplaySize(id string) (int, int) {
	cmd := exec.Command("adb", "-s", id, "shell", "wm ", "size")

	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Ekran boyut:", err)
		return 0, 0
	}
	outputStr := string(output)
	outputStr = outputStr[15:]
	//fmt.Println("Ekran Size : ", outputStr)
	x, _ := strconv.Atoi(outputStr[:4])
	a := outputStr[5:]
	y, _ := strconv.Atoi(a[:4])
	return x, y
}
