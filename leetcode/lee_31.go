package leetcode

func nextPermutation(nums []int)  {
	left, right, k := len(nums) - 2, len(nums) - 1, len(nums) - 1
	for left >= 0 && nums[left] > nums[right] {
		left--
		right--
	}
	if left >= 0 {
		for k > left && nums[k] <= nums[left] {
			k--
		}
		nums[left], nums[k] = nums[k], nums[left]
	}
	for i, j := right, len(nums) - 1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return
}