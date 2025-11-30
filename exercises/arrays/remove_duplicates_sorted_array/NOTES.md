# Solution Explanations & Decision Tree

## 🎯 Core Insight

**In a sorted array, all duplicates are adjacent!**
```
[1, 1, 2, 2, 2, 3] → Only compare with previous element
```

## 🔍 Decision Tree

```
Start: Remove duplicates from sorted array
            │
            ▼
    ┌───────────────┐
    │ Array sorted? │
    └───────────────┘
         │      │
       Yes      No → Sort first O(n log n) or use Hash Map O(n) space
         │
         ▼
    ┌────────────────┐
    │ In-place (O(1))│
    │ required?      │
    └────────────────┘
         │      │
       Yes      No → Create new array O(n) space
         │
         ▼
    ┌──────────────────┐
    │ Two Pointers     │
    │ Slow/Fast        │
    │ Time: O(n)       │
    │ Space: O(1)      │
    └──────────────────┘
```

## 💡 Solution: Two Pointers (Optimal)

### Strategy
- **Slow pointer (s)**: Tracks position for next unique element
- **Fast pointer (f)**: Explores array to find unique elements

### Algorithm
```
1. slow = 0 (first element always unique)
2. fast = 1 to n
3. If nums[fast] ≠ nums[slow]:
   → slow++
   → nums[slow] = nums[fast]
4. Return slow + 1
```

### Visual Example
```
nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]

Step 1: f=1, nums[1]=0 == nums[0]=0 → skip
Step 2: f=2, nums[2]=1 ≠ nums[0]=0 → unique! nums[1]=1
Step 3: f=3, nums[3]=1 == nums[1]=1 → skip
Step 4: f=4, nums[4]=1 == nums[1]=1 → skip
Step 5: f=5, nums[5]=2 ≠ nums[1]=1 → unique! nums[2]=2
Step 6: f=6, nums[6]=2 == nums[2]=2 → skip
Step 7: f=7, nums[7]=3 ≠ nums[2]=2 → unique! nums[3]=3
Step 8: f=8, nums[8]=3 == nums[3]=3 → skip
Step 9: f=9, nums[9]=4 ≠ nums[3]=3 → unique! nums[4]=4

Result: [0, 1, 2, 3, 4, ?, ?, ?, ?, ?]
Return: 5
```

### Complexity
- **Time:** O(n) - Single pass
- **Space:** O(1) - Two pointers only

## ⚠️ Common Pitfalls

1. **Both pointers start at 0** → First element gets skipped
2. **Assign before incrementing slow** → Overwrites current unique
3. **Return slow instead of slow+1** → Off-by-one error
4. **Forget empty array check** → Index out of bounds

## 🔑 Key Patterns

**Two Pointers - Slow/Fast Pattern:**
- Slow = write position
- Fast = read position  
- Overwrite in-place, no shifting needed

**Similar Problems:**
- Remove Element (#27)
- Move Zeroes (#283)
- Remove Duplicates II (#80)
