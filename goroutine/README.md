# Goroutine

## Coroutine과 CPS, CSP

> _Do not communicate by sharing memory; instead, share memory by communicating._
> — [Share Memory By Communicating](https://go.dev/blog/codelab-share)

Goroutine의 컨셉은 CSP(Communicating sequential processes)를 따른다.

## 성능

### Context Size

하드웨어인 프로세스 코어는 물리적 제한이 있다.
소프트웨어인 OS 스레드는 이를 효율적으로 활용하기 위해 Context Switching을 수행한다.

입출력을 예로 들어보자.
파일 입출력, 네트워크 입출력 등 읽고 쓰는 작업은 실제로 OS에서 처리한다.
프로세스에서 입출력 System call을 명령하면
프로세스는 OS가 입출력 작업을 끝낼 때까지 대기 상태가 된다.

이 때 OS 스레드가 가진 Context는 갖고 있는 내용이 많아서 사이즈가 크다.
하지만 경량 스레드(Light-weight thread)인 Goroutine은 Context의 사이즈가 작다.

실제로 Goroutine이 갖는 최소 Stack 사이즈는 2Kb다.

- [Go 1.2](https://go.dev/doc/go1.2#stacks): goroutine stack has been increased from 4Kb to 8Kb.
- [Go 1.4](https://go.dev/doc/go1.4#runtime): goroutine stack has decreased from 8Kb to 2Kb.

- [runtime/stack.go](https://github.com/golang/go/blob/go1.17.8/src/runtime/stack.go#L73)

```go
// The minimum size of stack used by Go code
_StackMin = 2048
```

Goroutine을 사용하면 OS 레이어에서 Context Switching이 발생하지 않는다.
Goroutine 만의 Context Switching이 발생하기 때문에 훨씬 가볍다.
Goroutine의 Context Switching? 메모리를 공유하기 때문에
Stack을 복사할 필요없이 Stack을 가리키는 포인터만 바꿔주면 된다

OS Thread는 프로세서에 따라 개수가 정해져있다.
Goroutine은 OS Thread의 리소스를 활용하는 것이기 때문에
리소스만 넉넉하다면 아주 많이 생성할 수 있다.

## 동시성 문제 (concurrency problems)

[예제 프로그램](./3mutex/main.go)과 같은 동시성 문제를 해결하기 위해 필요한 것이 Mutex (Mutual EXclusion)

Lock을 사용해서 메모리 자원을 하나의 Goroutine에서만 접근하게 만든다.

## 참고

- [Tucker의 Go 언어 프로그래밍](https://youtu.be/tRdODUXV3ik)
- [How Does the Goroutine Stack Size Evolve?](https://medium.com/a-journey-with-go/go-how-does-the-goroutine-stack-size-evolve-447fc02085e5)
