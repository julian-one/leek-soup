package easy

// 13. Roman to Integer
// Given a roman numeral, convert it to an integer.
//
// Time Complexity: O(n) - single pass through the string
// Space Complexity: O(1) - constant space for the map
func RomanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)

	for i := range n {
		current := roman[s[i]]

		// check not at end of string, current character is not smaller than next
		if i < n-1 && current < roman[s[i+1]] {
			total -= current
		} else {
			total += current
		}

	}
	return total
}
