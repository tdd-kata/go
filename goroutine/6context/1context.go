package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	//Context:
	//작업을 지시할 때
	//작업 가능 시간, 작업 취소 등의 조건을 지시할 수 있는
	//작업 명세서 역할.
	//프로세스나 API에 연계되는 모든 작업들이 하나의 맥락(context) 안에서 이뤄지도록 하는 것이다.
	//https://jaehue.github.io/post/how-to-use-golang-context/
	ctx, cancel := context.WithCancel(context.Background()) // (1) 컨텍스트 생성
	go PrintEverySecond(ctx)

	time.Sleep(5 * time.Second)
	cancel() // (2) 취소

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done(): // (3) 취소(cancel()) 확인
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
