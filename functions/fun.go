package main

import (
	"fmt"
	"sort"
)

func linearSearch(arr []int, ele int) (int, string) {
	if ele < 0 {
		return -1, "ele passed is negative"
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == ele {
			return i, "found at" 
		}
	}

	return -1, "ele not found"
}

func binarySearch(arr []int, ele int) (int, string) {
	if(ele < 0) {
		return -1, "element passed is negetive"
	}

	lb, ub := 0, len(arr) - 1;

	for lb <= ub {
		mid := lb + (ub - lb) / 2
		if arr[mid] == ele {
			return mid, "found at"
		} else if arr[mid] > ele {
			ub = mid - 1
		} else {
			lb = mid + 1
		}
	}

	return -1, "not found"
}

func main() {
	arr := []int{1, 2, 3, 5, 4}

	var ele int
	
	fmt.Scanf("%v", &ele)

	// idx, error := linearSearch(arr, ele)
	sort.Ints(arr)
	idx, error := binarySearch(arr, ele)

	if idx == -1 {
		fmt.Println(error)
	} else {
		fmt.Println(idx)
	}


}