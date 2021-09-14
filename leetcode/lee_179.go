package leetcode

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return strconv.Itoa(nums[i])+strconv.Itoa(nums[j]) > strconv.Itoa(nums[j])+strconv.Itoa(nums[i])
	})

	ans := ""
	for i := 0; i < len(nums); i++ {
		ans += strconv.Itoa(nums[i])
	}
	for len(ans) > 1 {
		if ans[0] == '0' {
			ans = ans[1:]
		} else {
			break
		}
	}
	return ans
}
