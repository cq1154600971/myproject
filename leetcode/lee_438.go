package leetcode

func findAnagrams(s string, p string) []int {
	n, m := len(s), len(p)
	ans := []int{}
	if n < m {
		return ans
	}
	pCnt := make([]int, 26)
	sCnt := make([]int, 26)
	for i := range p {
		pCnt[p[i] - 'a']++
	}

	left := 0
	for right := 0; right < n; right++ {
		curRight := s[right] - 'a'
		sCnt[curRight]++
		for sCnt[curRight] > pCnt[curRight] {
			curLeft := s[left] - 'a'
			sCnt[curLeft]--
			left++
		}
		if right - left + 1 == m {
			ans = append(ans, left)
		}
	}

	return ans
}