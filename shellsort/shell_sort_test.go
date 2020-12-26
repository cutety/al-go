package shellsort

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	nums := []int{3, 5, 2, 6, 4, 7, 1, 0}
	ShellSort(nums)
	fmt.Println(nums)
}

func ShellSort(nums []int) {
	l := len(nums)
	for step := l / 2; step >= 1; step /= 2 {
		for i := step; i < l; i += step {
			for j := i - step; j >= 0; j -= step {
				if nums[j+step] < nums[j] {
					nums[j], nums[j+step] = nums[j+step], nums[j]
					continue
				}
				break
			}
		}
	}
}
