package main

import "testing"

//go test -bench .
//Output:
//goos: windows
//goarch: amd64
//pkg: xpdojo.org/benchmark
//cpu: AMD Ryzen 5 3600X 6-Core Processor
//BenchmarkFibonacci1-12          578814240                2.058 ns/op
//BenchmarkFibonacci2-12          1000000000               0.2552 ns/op
// PASS
func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci1(1)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(1)
	}
}

func TestFibonacci1_1(t *testing.T) {
	if fibonacci1(1) != 1 {
		t.Error("Fibonacci(1) should be 1")
	}
}

func TestFibonacci1_2(t *testing.T) {
	if fibonacci1(2) != 1 {
		t.Error("Fibonacci(2) should be 1")
	}
}

func TestFibonacci1_3(t *testing.T) {
	if fibonacci1(3) != 2 {
		t.Error("Fibonacci(3) should be 2")
	}
}

func TestFibonacci1_4(t *testing.T) {
	if fibonacci1(4) != 3 {
		t.Error("Fibonacci(4) should be 3")
	}
}

func TestFibonacci1_5(t *testing.T) {
	if fibonacci1(5) != 5 {
		t.Error("Fibonacci(5) should be 5")
	}
}

func TestFibonacci1_6(t *testing.T) {
	if fibonacci1(6) != 8 {
		t.Error("Fibonacci(6) should be 8")
	}
}
