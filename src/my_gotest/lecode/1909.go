package lecode
//https://leetcode.cn/problems/remove-one-element-to-make-the-array-strictly-increasing/
func canBeIncreasing(nums []int) bool {
	notDecrease := 0

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] <= nums[i-1] {
			notDecrease++
			if i-2 >= 0 && i+1 < len(nums) && nums[i] <= nums[i-2] && nums[i+1] <= nums[i-1] {
				notDecrease++
				if notDecrease > 1 {
					return false
				}
			}

			if notDecrease > 1 {
				return false
			}
		}
	}

	return true
}

