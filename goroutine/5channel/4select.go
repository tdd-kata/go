package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Output:
// Square: 0
// Square: 4
// Square: 16
// Square: 36
// tick
// Square: 64
// Square: 100
// Square: 144
// Square: 196
// tick
// Square: 256
// Square: 324
// terminate
// main goroutine end
func main() {
	ch := make(chan int, 2)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	// close(ch)
	wg.Wait()

	fmt.Println("main goroutine end")
}

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)
	// select를 사용하면 여러 채널에서 값을 받을 수 있다.
	for {
		// case가 동시에 여러 개 해당되는 경우에는
		// 랜덤으로 하나만 선택된다.
		select {
		case <-tick:
			fmt.Println("tick")
		case <-terminate:
			fmt.Println("terminate")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}
