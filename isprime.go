package main

import "fmt"

func main() {
	var n int

	fmt.Println("enter number to check")
	fmt.Scan(&n)

	var prime = true
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			prime = false
			break
		}
	}

	if prime {
		fmt.Println("number is prime")
	} else {
		fmt.Println("number is not prime")
	}

}
