package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J")
	}

}

func main() {
	helpText := `Usage: clitimer minutes`

	var minutes int64

	if len(os.Args) != 2 {
		fmt.Println(helpText)
		return
	}

	minutes, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(helpText)
		return
	}

	timer := time.Minute * time.Duration(minutes)

	for timer > 0 {
		time.Sleep(time.Second)
		timer -= time.Second
		clearScreen()
		fmt.Printf("%.2d:%.2d", int(timer.Minutes()), int(timer.Seconds())%60)
	}
}
