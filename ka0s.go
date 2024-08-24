package main

import (
	"fmt"
	"runtime"
	"time"
)

const allocSize = 8 * 1024 * 1024 * 1024 // 8 GB
const cpuLoadIterations = 10000000

func cpuLoad() {
	for i := 0; i < cpuLoadIterations; i++ {
		// Busy work to consume CPU
	}
}

func main() {
	// Allocate 1 GB of memory
	mem := make([]byte, allocSize)
	fmt.Printf("Allocated %d MB of memory.\n", allocSize/(1024*1024))

	// Create a goroutine to consume CPU
	go cpuLoad()

	// Keep the program running to simulate high CPU load
	for {
		runtime.Gosched() // Yield the processor to allow other goroutines to run
		time.Sleep(1 * time.Second) // Sleep to simulate ongoing operation
	}
}
