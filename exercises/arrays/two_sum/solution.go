package twosum

// TwoSum finds two indices in nums whose values sum to target.
// Uses a hash map for O(n) time complexity.
// Time Complexity: O(n)
// Space Complexity: O(n)
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int) // value -> index

	for i, num := range nums {
		complement := target - num
		if j, exists := seen[complement]; exists {
			return []int{j, i}
		}
		seen[num] = i
	}

	return nil // No solution found
}

// TwoSumBruteForce is the naive O(n²) approach.
// Time Complexity: O(n²)
// Space Complexity: O(1)
func TwoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// TwoSumTwoPass uses two passes: first to build the map, then to find the complement.
// Time Complexity: O(n)
// Space Complexity: O(n)
func TwoSumTwoPass(nums []int, target int) []int {
	numMap := make(map[int]int) // value -> index

	// First pass: build the hash map
	for i, num := range nums {
		numMap[num] = i
	}

	// Second pass: find the complement
	for i, num := range nums {
		complement := target - num
		if j, exists := numMap[complement]; exists && j != i {
			return []int{i, j}
		}
	}

	return nil
}
