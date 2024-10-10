package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello"
		fmt.Println("Data has been sent to the channel")
	}()

	// data := <-channel
	// fmt.Println("Data received from the channel: ", data)
	time.Sleep(3 * time.Second)
	fmt.Println("Waiting for data from the channel")

}
