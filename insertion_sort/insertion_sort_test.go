package insertion_sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	list := []int{1,3,5,2,4}
	InsertionSort(list)
	fmt.Println(list)
}

func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		j := i - 1
		if cur < nums[j] {
			// 需要排序
			for ; j > 0 && cur < nums[j]; j-- {
				nums[j+1] = nums[j]
			}
			nums[j+1] = cur
		}
	}
}

// 1 3 5 2 4
// 1 3 5