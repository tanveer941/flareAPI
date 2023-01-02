package util

import (
	"sync"
	"time"
)

func Fetch3SecData(myChan chan string, wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	myChan <- "3 sec"
	defer wg.Done()
}

func Fetch5SecData(myChan chan string, wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	myChan <- "5 sec"
	defer wg.Done()
}
