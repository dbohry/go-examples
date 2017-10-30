package main

import (
	"fmt"
	"time"
)

func simulateEvent(name string, t time.Duration) {
	fmt.Println("Started ", name, ": Should take", t, "seconds.")
	time.Sleep(t * 1e9)
	fmt.Println("Finished ", name)
}

func main() {
	go simulateEvent("100m sprint", 10)
	go simulateEvent("Long jump", 6)
	go simulateEvent("High jump", 3)

	time.Sleep(10 * 1e9)
}