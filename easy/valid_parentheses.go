package easy

func isValid(s string) bool {
	var stack []string

	openParentheses := map[string]string{
		"]": "[",
		")": "(",
		"}": "{",
	}

	for _, r := range s {
		char := string(r)

		if char == "[" || char == "{" || char == "(" {
			stack = append(stack, char)
		}

		stackLen := len(stack) - 1

		if stackLen < 0 {
			return false
		}

		parentheses, foundParentheses := openParentheses[char]

		if foundParentheses {
			if parentheses != stack[stackLen] {
				return false
			}

			stack = stack[:stackLen]
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
