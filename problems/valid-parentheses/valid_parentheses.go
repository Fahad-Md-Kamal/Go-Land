package validparentheses

func ValidParenthesis(s *string) bool {
    var stack []string

	m := map[string]string{
		")" : "(",
		"}" : "{",
		"]" : "[",
	}

	for _, char := range *s {

		if mapVal, ok := m[string(char)]; ok {

			if len(stack) > 0{
				length := len(stack) - 1
				firstItem := stack[length]
				stack = stack[:length]
				if mapVal != firstItem{
					return false
				}
			}else{
			   stack = append(stack, string(char))
            }
		} else {
			stack = append(stack, string(char))
		} 
	}

	return len(stack) == 0
}
