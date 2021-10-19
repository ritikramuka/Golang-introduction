package main

import "fmt"

func main() {
	var n = 9
	fmt.Scan(&n)

	// var k = 8

	// var arr2 = int[k]

	var arr = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// for iteration
	for i := 0; i < n; i++ {
		fmt.Println(arr[i])
	}

	// foreach loop
	for _, v := range arr {
		fmt.Println(v)
	}

	i := 10
	for i > 0 {
		fmt.Println(i)
		i--
	}

}
