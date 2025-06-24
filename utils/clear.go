package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	cmd := exec.Command("clear") // Linux/Mac
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
