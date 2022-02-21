package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 2, 3}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		fmt.Println(left, mid, right)
		if nums[mid] != mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if left == len(nums)-1 && nums[left] == left {
		return left + 1
	} else {
		return left
	}
}
