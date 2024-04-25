package easy

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	r := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")

	s = r.ReplaceAllString(s, "")

	mayBePalindrome := make([]byte, 0)

	for i := len(s) - 1; i >= 0; i-- {
		mayBePalindrome = append(mayBePalindrome, s[i])
	}

	return s == string(mayBePalindrome)
}
