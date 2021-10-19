package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a, b := 0, 1

	for n > 0 {
		fmt.Println(a)
		c := a + b
		a = b
		b = c
		n--
	}
}
