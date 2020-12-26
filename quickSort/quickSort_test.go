package quicksort

import (
	"log"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{10, 16, 8, 12, 15, 6, 3, 9, 5}
	QuickSort(nums, 0, len(nums)-1)
	log.Println(nums)
}
