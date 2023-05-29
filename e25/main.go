package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	// время сейчас
	now := time.Now()
	// ждем, пока время с now не станет равным duration
	for time.Since(now) < duration {

	}
}

func main() {
	fmt.Println("Start and wait")
	sleep(5*time.Second)
	fmt.Println("End")
}