package fib

func main() {

}

func Fibonacci(i int) int {
	if i <= 2 {
		return 1
	}

	return Fibonacci(i - 2) + Fibonacci( i - 1)
}


