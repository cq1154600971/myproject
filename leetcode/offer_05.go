package leetcode

func ReplaceSpace(s string) string {
	var ans []byte
	for i := range s {
		if s[i] ==' ' {
			ans = append(ans, []byte("%20")...)
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
