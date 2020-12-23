package quickSort

func  Partition(nums []int, low, high int) int {
	pivot := nums[low]
	i := low
	j := high
	for i < j {
		if nums[i] >= pivot && nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else if nums[i] < pivot {
			i++
		} else if nums[j] >= pivot {
			j--
		}
	}
	nums[j], nums[low] = nums[low], nums[j]
	return j
}

func QuickSort(nums []int, low, high int) {
	if low < high {
		 j := Partition(nums, low, high)
		 QuickSort(nums, low, j)
		 QuickSort(nums, j + 1, high)
	}
}
