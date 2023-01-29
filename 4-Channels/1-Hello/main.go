package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string) // Empty

	// Thread 2
	go func() {
		channel <- "Hello world!" // Full
	}()

	// Thread 1
	msg := <-channel // Empties out
	fmt.Println(msg)
}
