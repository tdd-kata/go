package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

// go run 1dining.go
// B 스푼 획득
// A 포크 획득
// fatal error: all goroutines are asleep - deadlock!
func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(10)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", spoon, fork, "포크", "스푼")
	go diningProblem("B", fork, spoon, "스푼", "포크")
	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	count := 100
	for i := 0; i < count; i++ {
		first.Lock()
		fmt.Printf("%s %s 획득\n", name, firstName)
		second.Lock()
		fmt.Printf("%s %s 획득\n", name, secondName)

		fmt.Printf("%s 밥을 먹습니다\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}
