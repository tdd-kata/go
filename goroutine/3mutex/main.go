package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

// mutex.Lock()을 사용하지 않으면
// 단순히 1000을 더하고 빼는데
// panic: Balance should not be negative: -990
func (a *Account) DepositAndWithdraw() {
	mutex.Lock()
	defer mutex.Unlock()

	if a.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative: %d", a.Balance))
	}
	a.Balance += 1000
	time.Sleep(time.Millisecond)
	a.Balance -= 1000
}

func main() {
	var wg sync.WaitGroup

	// account := Account{Balance: 10}
	account := &Account{10}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				account.DepositAndWithdraw()
			}
		}()
	}
	wg.Wait()
}
