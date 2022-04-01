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
// Square: 64
// Square: 100
// Square: 144
// Square: 196
// Square: 256
// Square: 324
// main goroutine end
func main() {
	ch := make(chan int, 2)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	// 채널을 닫지 않아서 무한 대기하는 고루틴을
	// 'zombie goroutine' 또는 'goroutine leak'이라고 부른다.
	// 이런 경우는 채널을 닫아서 무한 대기하는 고루틴을 제거하는 것이 좋다.
	close(ch)

	// wg.Wait()을 사용하면 다른 고루틴이 종료될 때까지 대기할 수 있다.
	// deadlock!
	// range ch으로 계속 데이터가 들어오길 기다리는데
	// main에서도 채널을 닫지 않는다면 wg.Wait()으로 계속 대기하기 때문에
	// 모든 goroutine이 대기 상태가 된다.
	wg.Wait()

	fmt.Println("main goroutine end")
}

func square(wg *sync.WaitGroup, ch chan int) {
	// channel에 range를 걸면 channel이 닫히는 것을 막을 수 있다.
	for n := range ch {
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}

	wg.Done()
}
