package main

import (
	"fmt"
	"runtime"
	"time"
)

const allocSize = 6 * 1024 * 1024 * 1024 // 6 GB
const cpuLoadIterations = 1000000 // Adjust this based on CPU load preference

func cpuLoad() {
	for i := 0; i < cpuLoadIterations; i++ {
		// Busy work to consume CPU
	}
}

func main() {
	// Allocate 6 GB of memory
	mem := make([]byte, allocSize)
	fmt.Printf("Allocated %d GB of memory.\n", allocSize/(1024*1024*1024))

	// Use the allocated memory to prevent unused variable error
	_ = mem

	// Create a goroutine to consume CPU
	go cpuLoad()

	// Keep the program running to simulate high CPU load
	for {
		runtime.Gosched() // Yield the processor to allow other goroutines to run
		time.Sleep(1 * time.Second) // Sleep to simulate ongoing operation
	}
}
