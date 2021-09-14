package leetcode

// func lengthOfLongestSubstring(s string) int {
//     cnt := map[byte]int{}
//     rk, ans := 0, 0
//     n := len(s)
//     for i := 0; i < n; i++ {
//         if i != 0 {
//             delete(cnt, s[i - 1])
//         }
//         for rk < n && cnt[s[rk]] == 0 {
//             cnt[s[rk]]++
//             rk++
//         }
//         ans = max(ans, rk - i)
//     }
//     return ans
// }
func lengthOfLongestSubstring(s string) int {
	cnt := map[byte]int{}
	ans, left := 0, 0
	for i := 0; i < len(s); i++ {
		if _, ok := cnt[s[i]]; ok {
			left = max(left, cnt[s[i]])
		}
		cnt[s[i]] = i + 1
		ans = max(ans, i + 1 - left)
	}
	return ans
}
//func max(x, y int) int {
//	if x < y {
//		return y
//	}
//	return x
//}