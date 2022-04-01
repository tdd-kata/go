package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	ch <- 9
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	n := <-ch
	// time.Sleep(time.Second)
	fmt.Printf("Square: %d\n", n*n)
}
