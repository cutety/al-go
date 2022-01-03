package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	count := search(nums, target)
	fmt.Println(count == 0)

	nums2 := []int{5, 7, 7, 8, 8, 10}
	target2 := 8
	fmt.Println(search(nums2, target2) == 2)

	nums3 := []int{2, 2}
	target3 := 2
	fmt.Println(search(nums3, target3) == 2)

	nums4 := []int{1}
	target4 := 1
	fmt.Println(search(nums4, target4) == 1)

	nums5 := []int{1}
	target5 := 0
	fmt.Println(search(nums5, target5) == 0)

	nums6 := []int{1, 2, 3}
	target6 := 1
	fmt.Println(search(nums6, target6) == 1)

	nums7 := []int{3, 3, 3}
	target7 := 3
	fmt.Println(search(nums7, target7) == 3)

	fmt.Println(search([]int{0, 0, 1, 2, 2}, 2) == 2)

}

func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	mid := (left + right - 1) / 2
	if left < right {
		if nums[mid] == target {
			var leftCount, rightCount int
			if mid > left && nums[mid-1] == target {
				leftCount = search(nums[left:mid], target)
			}
			if mid < right-1 && nums[mid+1] == target {
				rightCount = search(nums[mid+1:right], target)
			}
			return leftCount + rightCount + 1
		} else if nums[mid] < target && mid < right {
			return search(nums[mid+1:right], target)
		} else if nums[mid] > target && mid > left {
			return search(nums[left:mid], target)
		}
	}

	return 0
}
