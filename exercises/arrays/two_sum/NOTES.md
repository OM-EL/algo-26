# Problem-Solving Notes for Two Sum

## Initial Thoughts

When first reading the problem:
- We need to find TWO numbers that add up to a target
- We need to return their INDICES, not the values themselves
- Each input has exactly one solution (guaranteed)
- Cannot use the same element twice

### Key Questions to Ask
1. Is the array sorted? (No)
2. Can there be negative numbers? (Yes)
3. Can there be duplicate values? (Yes, see Example 3)
4. What should we return if no solution exists? (Not applicable per constraints)

## Approach 1: Brute Force

### Idea
Check every possible pair of numbers in the array.

### Algorithm
1. For each element at index i (from 0 to n-2)
2. For each element at index j (from i+1 to n-1)
3. If nums[i] + nums[j] == target, return [i, j]

### Complexity Analysis
- **Time Complexity:** O(n²) - Two nested loops
- **Space Complexity:** O(1) - No extra space used

### Pros/Cons
- ✅ Simple and intuitive
- ✅ No extra space needed
- ❌ Very slow for large arrays
- ❌ Does not meet the follow-up requirement

## Approach 2: Hash Map (One Pass)

### Idea
For each number, we're looking for its complement (target - num). Instead of searching through the array, we can use a hash map to check in O(1) time.

### Key Insight
As we iterate through the array, for each element `num`, we need to find if `target - num` exists in the array. If we store previously seen elements in a hash map, we can check this in O(1) time.

**The Magic Equation:** If `a + b = target`, then `b = target - a`

### Algorithm
1. Create an empty hash map (value -> index)
2. For each element at index i:
   - Calculate complement = target - nums[i]
   - If complement exists in hash map, return [hashmap[complement], i]
   - Otherwise, add nums[i] -> i to hash map
3. Return nil (no solution found)

### Why It Works
```
nums = [2, 7, 11, 15], target = 9

Step 1: num = 2
  complement = 9 - 2 = 7
  7 not in map
  map = {2: 0}

Step 2: num = 7
  complement = 9 - 7 = 2
  2 IS in map at index 0!
  Return [0, 1] ✓
```

### Complexity Analysis
- **Time Complexity:** O(n) - Single pass through array, O(1) hash map operations
- **Space Complexity:** O(n) - Hash map can store up to n elements

### Pros/Cons
- ✅ Optimal time complexity
- ✅ Meets the follow-up requirement
- ❌ Uses extra space
- ❌ Slightly more complex to understand

## Approach 3: Hash Map (Two Pass)

### Idea
First build a complete hash map, then search for complements.

### Algorithm
1. First pass: Build hash map of all values -> indices
2. Second pass: For each element, check if complement exists
3. Make sure complement isn't the same element (check index)

### Complexity Analysis
- **Time Complexity:** O(n) - Two separate passes
- **Space Complexity:** O(n) - Hash map stores all elements

### Note
This is slightly less efficient than one-pass since we go through the array twice, but it's easier to understand and implement correctly.

## Edge Cases to Consider

- [x] Array with only two elements [3, 3] -> [0, 1]
- [x] Negative numbers [-1, -2, -3], target = -3 -> [0, 1]
- [x] Target is zero [0, 4, 3, 0] -> [0, 3]
- [x] Large numbers near int limits
- [x] Solution is at the end of array

## Visual Representation

```
Array:  [2, 7, 11, 15]
Index:   0  1   2   3
Target: 9

Looking for pairs:
2 + ? = 9  →  ? = 7 ✓ (found at index 1)

Answer: [0, 1]
```

## Lessons Learned

1. **Hash maps trade space for time** - When you need O(1) lookup, think hash map
2. **Complement technique** - Instead of looking for two numbers, look for one number and its complement
3. **Order matters in one-pass** - We find the second element, so we return [earlier_index, current_index]

## Related Problems

- [15. 3Sum](https://leetcode.com/problems/3sum/) - Extension to three numbers
- [18. 4Sum](https://leetcode.com/problems/4sum/) - Extension to four numbers
- [167. Two Sum II](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) - Sorted array version
- [653. Two Sum IV](https://leetcode.com/problems/two-sum-iv-input-is-a-bst/) - BST version
