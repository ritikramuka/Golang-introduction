package main

import "fmt"

func main() {
	var n int

	fmt.Print("size of array: ")
	fmt.Scan(&n)

	var arr = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	var max = arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	for i := max; i > 0; i-- {
		for _, v := range arr {
			if v >= i {
				fmt.Print("*\t")
			} else {
				fmt.Print("\t")
			}
		}
		fmt.Println()
	}

}
