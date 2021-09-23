package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

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
		fmt.Printf("%.2d:%.2d\r", int(timer.Minutes()), int(timer.Seconds())%60)
	}
	fmt.Println("Finish")
}
