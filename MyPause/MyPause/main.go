package main

import (
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	delay := 10 //默认等待10秒
	if len(os.Args) > 1 {
		delay = ParseDelaySeconds(os.Args[1], delay)
	}
	time.Sleep(time.Duration(delay) * time.Second)
}

func ParseDelaySeconds(delayStr string, defaultDelay int) int {
	reg := regexp.MustCompile("[^0-9]")
	delayPara := reg.ReplaceAllString("0"+delayStr, "")
	newDelay, _ := strconv.Atoi(delayPara)
	if newDelay <= 0 {
		newDelay = defaultDelay
	}
	return newDelay
}
