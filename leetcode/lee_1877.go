package leetcode

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := 0
	for i, num := range nums[:n / 2] {
		ans = max(ans, num + nums[n - 1 - i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}