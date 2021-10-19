package main

import "fmt"

func fibo(n int) int {

	if n <= 1 {
		return n
	}

	nm1 := fibo(n - 1)
	nm2 := fibo(n - 2)

	return nm1 + nm2
}

func main() {
	var n int
	fmt.Scan(&n)

	ans := fibo(n)
	fmt.Println(ans)
}
