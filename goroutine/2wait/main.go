package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func CoSumAtoB(start, end int) {
	sum := SumAtoB(start, end)
	fmt.Printf("%d부터 %d까지 합계는 %d입니다\n", start, end, sum)
	wg.Done() // 작업이 완료될 때마다 호출
}

func SumAtoB(start, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}

//go run main.go
//1부터 100000000까지 합계는 5000000050000000입니다
//1부터 100000000까지 합계는 5000000050000000입니다
//1부터 100000000까지 합계는 5000000050000000입니다
//...
func main() {
	worker := 1000
	wg.Add(worker) // 작업 개수 설정
	for i := 0; i < worker; i++ {
		go CoSumAtoB(1, 100_000_000)
	}
	wg.Wait() // 모든 작업이 완료될 때까지 대기
}
