package medium

func LengthOfLongestSubstring(s string) int {
	mx := 0

	start := 0
	end := 0

	mp := map[byte]int{}

	for end = 0; end < len(s); end++ {
		mp[s[end]]++
		for mp[s[end]] > 1 {
			key := s[start]
			mp[key]--
			start++
		}
		mx = max(mx, end-start+1)
	}

	return mx
}
