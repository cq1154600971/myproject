package leetcode

import "sort"

func MaxProduct(words []string) int {
	sort.Slice(words, func(i , j int) bool {
		return len(words[i]) > len(words[j])
	})
	nums := make([]int, len(words))
	cnts := make([]int, len(words))
	n := len(words)
	for i := 0; i < n; i++ {
		nums[i] = len(words[i])
		tmp := 0
		for j := 0; j < len(words[i]); j++ {
			tmp |= 1 << int(words[i][j] - 'a')
		}
		cnts[i] = tmp
	}
	ans := 0
	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] * nums[j] <= ans {
				break
			}
			if cnts[i] & cnts[j] == 0 {
				ans = nums[i] * nums[j]
				break
			}
		}
	}
	return ans
}
