# TDD로 배우는 Go

- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

## 초기화

```bash
go mod init xpdojo.org/learn-go-with-tests
```

## 테스트 실행

```shell
go test ./...
# ok      xpdojo.org/learn-go-with-tests/integers/v1      0.172s
```

## 벤치마크

```shell
go test -bench=./...
```

## 테스트 커버리지

```shell
go test -cover ./...
```