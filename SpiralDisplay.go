package main

import "fmt"

func main() {
	var n int
	var m int

	fmt.Print("size of array: ")
	fmt.Scan(&n)
	fmt.Scan(&m)

	var arr = [][]int{}

	for i := 0; i < n; i++ {
		var ar = []int{}
		for j := 0; j < m; j++ {
			var d int
			fmt.Scan(&d)
			ar = append(ar, d)
		}

		arr = append(arr, ar)
	}

	// int minr=0, minc=0, maxr=n-1, maxc=m-1
	minr, minc, maxr, maxc := 0, 0, n-1, m-1

	var rem = n * m

	for rem > 0 {
		for r := minr; r <= maxr; r++ {
			fmt.Println(arr[r][minc])
			rem--
		}

		for c := minc + 1; c <= maxc; c++ {
			fmt.Println(arr[maxr][c])
			rem--
		}

		if rem == 0 {
			break
		}
		for r := maxr - 1; r >= minr; r-- {
			fmt.Println(arr[r][maxc])
			rem--
		}

		for c := maxc - 1; c > minc; c-- {
			fmt.Println(arr[minr][c])
			rem--
		}
		minr++
		minc++
		maxr--
		maxc--
	}
	fmt.Println()
}
