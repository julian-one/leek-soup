package easy

// 1. Two Sum
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
//
// Time Complexity: O(n) - single pass through the array
// Space Complexity: O(n) - hash map to store values and indices
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	// Example: nums=[3,8,4,1], target=7
	for i, num := range nums {
		// What number do we need to reach the target?
		// num=3, complement=7-3=4 (need 4 to make 7)
		// num=8, complement=7-8=-1 (need -1 to make 7)
		// num=4, complement=7-4=3 (need 3 to make 7)
		complement := target - num

		// Have we seen the complement before?
		// seen[4]? No
		// seen[-1]? No
		// seen[3]? Yes! at index 0
		if j, exists := seen[complement]; exists {
			return []int{j, i} // Return [0,2]
		}

		// Store current number and its index
		// seen[3] = 0
		// seen[8] = 1
		// seen[4] = 2
		seen[num] = i
	}

	return nil
}
