package easy

// 14. Longest Common Prefix
// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
//
// Time Complexity: O(S) - where S is the sum of all characters in all strings
// Space Complexity: O(1) - only using constant extra space
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i] // Get byte at position i

		for j := 1; j < len(strs); j++ {
			// Ensure we haven't exceeded the length of the string
			// OR the characters don't match
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}

	}
	return strs[0]
}
