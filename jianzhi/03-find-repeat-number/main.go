package main
import "fmt"

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	// 1,3,2,0,2,5,3
	// 3,1,2,0,2,5,3
	// 0,1,2,3,2,5,3
	// 0,1,2,3, (2)

 fmt.Println(findRepeatNumber(nums))

}

func findRepeatNumber(nums []int) int {
	i := 0

	for {
		if nums[i] != i {
			if nums[i] != nums[nums[i]] {
				nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
			} else {
				return nums[i]
			}	
		} else {
			i++
		}
	}
}