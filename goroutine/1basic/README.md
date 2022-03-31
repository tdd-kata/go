# Simple Goroutine

```bash
go mod init xpdojo.org/basic
go run main.go
# go build
# ./basic
```

```mermaid
sequenceDiagram
    participant PrintHangul()
    participant PrintNumber()
    participant main()
    Note right of main(): 3,000ms
    PrintHangul()-->>PrintNumber(): 가 (300ms)
    PrintNumber()-->>PrintHangul(): 1 (400ms)
    PrintHangul()-->>PrintNumber(): 나 (600ms)
    PrintNumber()-->>PrintHangul(): 2 (800ms)
    PrintHangul()-->>PrintNumber(): 다 (900ms)
    PrintHangul()-->>PrintNumber(): 라 (1,200ms)
    PrintNumber()-->>PrintHangul(): 3 (1,200ms)
    PrintHangul()-->>PrintNumber(): 마 (1,500ms)
    PrintNumber()-->>PrintHangul(): 4 (1,600ms)
    PrintHangul()-->>PrintNumber(): 바 (1,800ms)
    PrintNumber()-->>PrintHangul(): 5 (2,000ms)
    PrintHangul()-->>PrintNumber(): 사 (2,100ms)
```

[mermaid.js](http://mermaid-js.github.io/)
