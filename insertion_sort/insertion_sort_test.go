package insertion_sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	list := []int{5,9,1,6,8,14,6,49,25,4,6,3}
	InsertionSort(list)
	fmt.Println(list)
}

func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		j := i - 1
		if cur < nums[j] {
			// 需要排序
			for ; j >= 0 && cur < nums[j]; j-- {
				nums[j+1] = nums[j]
			}
			nums[j+1] = cur
		}
	}
}

// output [1 3 4 5 6 6 6 8 9 14 25 49]

