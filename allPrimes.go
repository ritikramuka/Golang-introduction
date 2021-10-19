package main

import "fmt"

func isPrime(n int) bool {
	var prime = true

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			prime = false
			break
		}
	}

	return prime
}

func main() {
	var n int

	fmt.Println("enter number")
	fmt.Scan(&n)

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}

}
