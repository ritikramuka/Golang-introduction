package main

import "fmt"

func prepareArray(n int, m int) [][]int {
	arr := make([][]int, n);

	for i:=0; i < n; i++ {
		arr[i] = make([]int, m)
	}

	return arr
}

func input(arr [][]int) {
	n := len(arr)
	m := len(arr[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
}

func main()  {
	var n,m int 
	fmt.Scan(&n)
	fmt.Scan(&m)

	arr := prepareArray(n, m)
	input(arr)

	sr, sc := 0, 0
	dir := make([]int)

	for true {
		if sr + 
	}
}