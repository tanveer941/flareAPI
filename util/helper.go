package util

import "time"

func Fetch3SecData(myChan chan string) {
	time.Sleep(3 * time.Second)
	myChan <- "3 sec"
}

func Fetch5SecData(myChan chan string) {
	time.Sleep(5 * time.Second)
	myChan <- "5 sec"
}
