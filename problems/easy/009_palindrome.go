package easy

// 9. Palindrome Number
// Given an integer x, return true if x is a palindrome, and false otherwise.
//
// Time Complexity: O(log n) - we process each digit once
// Space Complexity: O(1) - only using constant extra space
func IsPalindrome(x int) bool {
	// Negatives are not palindromes
	if x < 0 {
		return false
	}

	original, reversed := x, 0

	for x > 0 {
		// Get the last most digit using modulo
		// 123 % 10 = 3
		current := x % 10

		// Append the reversed shift left and add current digit
		// reversed=0, current=3: 0*10 + 3 = 3
		// reversed=3, current=2: 3*10 + 2 = 32
		// reversed=32, current=1: 32*10 + 1 = 321

		reversed = reversed*10 + current
		// Remove the last digit from x
		// 123 / 10 = 12
		// 12 / 10 = 1
		// 1 / 10 = 0
		x /= 10
	}
	return original == reversed
}
