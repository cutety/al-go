package main

import (
	"fmt"
)

// 2,3,4,0,1 left:0, right: 4, mid:2
// 4 0 1 left:0, right: 2, mid:1
// 0 1 left:0 , right:1, mid: 0
//

func main() {
	fmt.Println(minArray([]int{3, 1}))
}

func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		mid := left + (right-left)/2
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
		} else {
			right--
		}
	}

	return numbers[left]
}
