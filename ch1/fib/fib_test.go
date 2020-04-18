package fib

import "testing"

func TestFibonacci(t *testing.T) {
	t.Run("fib 5", func(t *testing.T) {
		got := Fibonacci(5)
		want := 5

		assertTest(t, got, want)
	})

	t.Run("fib 10", func(t *testing.T) {
		got := Fibonacci(9)
		want := 34

		assertTest(t, got, want)
	})

	t.Run("fib 1", func(t *testing.T) {
		got := Fibonacci(1)
		want := 1

		assertTest(t, got, want)
	})

	t.Run("fib 2", func(t *testing.T) {
		got := Fibonacci(2)
		want := 1

		assertTest(t, got, want)
	})
}

func assertTest(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}

