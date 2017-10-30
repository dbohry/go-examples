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
	go simulateEvent("Example 1", 10)
	go simulateEvent("Example 2", 3)
	go simulateEvent("Example 3", 5)
	go simulateEvent("Example 4", 12)

	time.Sleep(13 * 1e9)
}