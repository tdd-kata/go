# Benchmark

성능 검사 시 사용한다.

## 표현 규약

- 파일명이 `_test.go`로 끝나야 한다.
- testing 패키지를 import 해야 한다.
- 벤치마크 코드는 `func BenchmarkXxxx(b *testing.B)` 형태여야 한다.
- 실행 시간을 측정하기 위해서는 `b.N` 을 사용해야 한다.

