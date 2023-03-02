package algorithms

func binary_search(nums []int, target int) (int) {
	low := 0
	hight := len(nums)

	for low <= hight {
		mid := (low + hight) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hight = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		}
	}

	return -1
}