package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go square()
	ch <- 9
	fmt.Println("Never print")
}

func square() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Printf("Sleep\n")
	}
}
