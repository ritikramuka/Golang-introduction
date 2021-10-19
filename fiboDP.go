package main

import "fmt"

var memo = []int{}

func fibo(n int) int {

	if n <= 1 {
		return n
	}

	if memo[n] != 0 {
		return memo[n]
	}

	nm1 := fibo(n - 1)
	nm2 := fibo(n - 2)

	memo[n] = nm1 + nm2
	return nm1 + nm2
}

func main() {
	var n int
	fmt.Scan(&n)

	memo = make([]int, n)

	ans := fibo(n)
	fmt.Println(ans)
}
