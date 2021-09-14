package leetcode

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	return check(s, 0, len(s) - 1)
}
func check(s string, i, j int) bool {
	for i < j {
		for ; !unicode.IsDigit(rune(s[i])) && !unicode.IsLetter(rune(s[i])) && i <= j; i++ {
		}
		for ; (s[j] < 'a' || s[j] > 'z'|| s[i] < '0' || s[i] > '9') && i <= j; j-- {
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}