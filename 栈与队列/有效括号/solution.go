package main

func isValid(s string) bool {
	cs := []byte(s)
	stack := []byte{}
	pairs := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	for i := 0; i < len(cs); i++ {
		le, ok := pairs[cs[i]]
		if ok {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != le {
				return false
			} else {
				stack = stack[:len(stack)-1]

			}
		} else {
			stack = append(stack, cs[i])
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
