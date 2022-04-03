# Goroutine

## Coroutine과 CSP

> _Do not communicate by sharing memory; instead, share memory by communicating._
> — [Share Memory By Communicating](https://go.dev/blog/codelab-share)

Goroutine의 컨셉은 CSP(Communicating sequential processes)를 따른다.
대부분의 프로그램은 스레드끼리 공유할 데이터가 있으면 메모리에 적재된 동시성 객체를 공유했다.
하지만 이는 힘든 메모리 매니지먼트, 알 수 없는 예외, 데드락 등을 수반한다.
그래서 대규모 프로그램이나 복잡한 프로그램에서 올바르게 활용하기 어렵다.
대신 CSP는 채널 간 전달(커뮤니케이션)을 통해 메모리를 공유한다.
Go 언어에서 채널(`chan`)은 스레드의 안전성을 제공하는 큐(queue)다.

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
예를 들어, Java의 스레드는 약 1Mb 메모리가 필요하고 스레드 간 여유 공간인 guard page까지 필요하다.
이에 반해 Goroutine이 갖는 최소 Stack 사이즈는 2Kb다.

- [Go 1.2](https://go.dev/doc/go1.2#stacks): goroutine stack has been increased from 4Kb to 8Kb.
- [Go 1.4](https://go.dev/doc/go1.4#runtime): goroutine stack has decreased from 8Kb to 2Kb.

- [runtime/stack.go](https://github.com/golang/go/blob/go1.17.8/src/runtime/stack.go#L73)

```go
// The minimum size of stack used by Go code
_StackMin = 2048
```

또한 Context Switching 시 일반적인 스레드는
16개의 범용 레지스터, PC, SP, segment 레지스터, 16개의 XMM 레지스터 등 많은 것을 저장하고 복구해야 하지만
Goroutine은 3개의 레지스터(PC, SP, DX)만 저장 및 복구하면 되기 때문에 Context Switching 비용이 매우 적다.

Goroutine을 사용하면 OS 레이어에서 Context Switching이 발생하지 않는다.
Goroutine 만의 Context Switching이 발생하기 때문에 훨씬 가볍다.
Goroutine의 Context Switching? 메모리를 공유하기 때문에
Stack을 복사할 필요없이 Stack을 가리키는 포인터만 바꿔주면 된다.

OS Thread는 프로세서에 따라 개수가 정해져있다.
Goroutine은 OS Thread의 리소스를 활용하는 것이기 때문에
리소스만 넉넉하다면 아주 많이 생성할 수 있다.

## 동시성 문제 (concurrency problems)

[예제 프로그램](./3mutex/main.go)과 같은 동시성 문제를 해결하기 위해 필요한 것이 Mutex (Mutual EXclusion)

Lock을 사용해서 메모리 자원을 하나의 Goroutine에서만 접근하게 만든다.

## 참고

- <7가지 동시성 모델> 6장. 순차 프로세스 통신 - 임백준
- [고(golang) 언어에서의 동시성 모델](https://okky.kr/article/445939)
  - [원문: Concurrency in Golang](http://www.minaandrawos.com/2015/12/06/concurrency-in-golang/) - Mina Andrawos
- [Tucker의 Go 언어 프로그래밍](https://youtu.be/tRdODUXV3ik)
  - [tuckersGo/musthaveGo](https://github.com/tuckersGo/musthaveGo)
- [How Does the Goroutine Stack Size Evolve?](https://medium.com/a-journey-with-go/go-how-does-the-goroutine-stack-size-evolve-447fc02085e5)
- [고루틴(goroutine) Deep Dive](https://mainfunction.tistory.com/entry/GoGolang-%EA%B3%A0%EB%A3%A8%ED%8B%B4goroutine-Deep-Dive)
