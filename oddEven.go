package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i += 2 {
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Println()

	for i := 2; i <= n; i += 2 {
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Println()

}