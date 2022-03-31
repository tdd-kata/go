package main

import (
	"fmt"
	"time"
)

// go run main.go
func main() {
	go PrintHangul()
	go PrintNumbers()

	time.Sleep(3 * time.Second)
	// 가 1 나 2 다 3 라 마 4 바 5 사
}

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
