package easy

// 20. Valid Parentheses (Stack)
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
//
// Time Complexity: O(n) - single pass through the string
// Space Complexity: O(n) - stack can hold at most n/2 elements
func IsValid(s string) bool {
	stack := []byte{}
	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	for i := range len(s) {
		char := s[i] // byte conversion
		if opening, isClosing := pairs[char]; isClosing {
			// Must match top of stack
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			stack = stack[:len(stack)-1] // pop matching
		} else {
			stack = append(stack, char) // push
		}
	}
	// stack must be empty to be valid
	return len(stack) == 0
}
