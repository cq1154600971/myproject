package leetcode

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	res :=[]int{}
	sort.Ints(candidates)
	var dfs func(target, k int)
	dfs = func(target, k int){
		if target == 0 {
			ans = append(ans, append([]int{}, res...))
			return
		}
		for i := k; i < len(candidates); i++ {
			if target - candidates[i] < 0 {
				break
			}
			res = append(res, candidates[i])
			dfs(target - candidates[i], i)
			res = res[:len(res) - 1]
		}
	}
	dfs(target, 0)
	return ans
}