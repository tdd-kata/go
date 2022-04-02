package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	//Context Wrapping
	//한번 생성된 컨텍스트는 변경할 수 없다.
	//그래서 컨텍스트에 값을 추가하고 싶을 때는 context.WithValue() 함수로 새로운 컨텍스트를 만들어 주어야 한다.
	//https://velog.io/@kimmachinegun/Go-context.WithValue-%EC%95%88%EC%A0%84%ED%95%98%EA%B2%8C-%EC%82%AC%EC%9A%A9%ED%95%98%EA%B8%B0
	ctx := context.WithValue(context.Background(), "number", 9) // (1) 컨텍스트에 값을 추가한다
	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil { // (2) 컨텍스트에서 값을 읽어온다.
		n := v.(int)
		fmt.Printf("Square:%d", n*n)
	}
	wg.Done()
}
