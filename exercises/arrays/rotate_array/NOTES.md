# Solution Explanations & Decision Tree

## 🎯 Core Insight

**Reverse pattern: Reverse entire array, then reverse each part!**
```
[1,2,3,4,5,6,7], k=3
→ Reverse all: [7,6,5,4,3,2,1]
→ Reverse [0,k): [5,6,7,4,3,2,1]
→ Reverse [k,n): [5,6,7,1,2,3,4] ✓
```

**Why this works:** Reversing twice brings elements to their final position.

## 🔍 Decision Tree

```
Start: Rotate array by k positions
            │
            ▼
    ┌───────────────┐
    │ Extra space   │
    │ allowed?      │
    └───────────────┘
         │      │
        Yes     No → Must be in-place
         │           │
         │           ▼
         │      ┌──────────────────┐
         │      │ Three Reversals  │
         │      │ Time: O(n)       │
         │      │ Space: O(1) ✓    │
         │      └──────────────────┘
         │
         ▼
    ┌──────────────────┐
    │ Copy to new array│
    │ Time: O(n)       │
    │ Space: O(n)      │
    └──────────────────┘
```

## 💭 How to Think About This

**First Observations:**
1. "Rotate right by k" - last k elements move to front
2. k can be > n - need to handle with modulo
3. In-place requirement - can't use extra array

**Questions to Ask Yourself:**
- Q: What happens when k > array length?
- A: Use k % n (rotating by n is same as no rotation)

- Q: Where do elements end up after rotation?
- A: Element at index i goes to index (i + k) % n

- Q: Can I avoid moving elements one by one?
- A: Yes! Use reversal trick - elegant and O(1) space

**Key Realization:**
```
[1,2,3,4,5,6,7], k=3

Visual breakdown:
Original:    [1,2,3,4 | 5,6,7]
             └─part2─┘ └part1┘

After rotate: [5,6,7 | 1,2,3,4]
              └part1┘ └─part2─┘

Observation: Parts swap positions AND reverse internally!
```

**Pattern Recognition:**
When you need to "move chunks" in-place → Think reversal!

## 💡 Solution Approach 1: Three Reversals (Optimal)

### Strategy (Plain Language)
Instead of moving elements k times, use three reversals:
1. Reverse entire array
2. Reverse first k elements
3. Reverse remaining n-k elements

### Why This Works
**Mathematical insight:**
```
Let A = first n-k elements, B = last k elements
Original: [A][B]
Goal:     [B][A]

Step 1: Reverse all → [B'][A']  (reversed versions)
Step 2: Reverse B'  → [B][A']
Step 3: Reverse A'  → [B][A] ✓
```

**Intuition:**
Reversing "undoes" the incorrect order from the first reversal.

### Algorithm
```
1. k = k % n  (handle k > n)
2. Reverse nums[0...n-1]     // Entire array
3. Reverse nums[0...k-1]     // First k elements
4. Reverse nums[k...n-1]     // Remaining elements
```

### Visual Example
```
nums = [1,2,3,4,5,6,7], k = 3

Step 0: Original
[1, 2, 3, 4, 5, 6, 7]

Step 1: Reverse entire array
[7, 6, 5, 4, 3, 2, 1]
 └──────────────────┘

Step 2: Reverse first k=3 elements
[5, 6, 7, 4, 3, 2, 1]
 └─────┘

Step 3: Reverse remaining n-k=4 elements
[5, 6, 7, 1, 2, 3, 4]
          └────────┘

Result: [5,6,7,1,2,3,4] ✓
```

## 🧮 Why This Works (Mathematical Proof)

**Claim:** Three reversals correctly rotate array by k positions.

**Proof:**
Let array be divided as: `nums = [A₁, A₂, ..., Aₙ₋ₖ | B₁, B₂, ..., Bₖ]`

After Step 1 (reverse all):
```
[Bₖ, Bₖ₋₁, ..., B₁ | Aₙ₋ₖ, ..., A₂, A₁]
```

After Step 2 (reverse first k):
```
[B₁, B₂, ..., Bₖ | Aₙ₋ₖ, ..., A₂, A₁]
```

After Step 3 (reverse last n-k):
```
[B₁, B₂, ..., Bₖ | A₁, A₂, ..., Aₙ₋ₖ]
```

This is exactly the rotated array! ✓

**Why three reversals specifically?**
- 1 reversal: Wrong order
- 2 reversals: Partial fix
- 3 reversals: Complete fix (brings everything to correct position)

## 💡 Solution Approach 2: Cyclic Replacements

### Strategy
Follow cycles: each element displaces the next in a circular pattern.

### Algorithm
```
1. count = 0
2. For start from 0 to n-1:
   - current = start
   - prev = nums[start]
   - Do:
     - next = (current + k) % n
     - temp = nums[next]
     - nums[next] = prev
     - prev = temp
     - current = next
     - count++
   - While current != start
   - If count == n, break
```

### Complexity
- **Time:** O(n) - Each element moved once
- **Space:** O(1)

**Note:** More complex, harder to implement correctly.

## 💡 Solution Approach 3: Extra Array (Not In-Place)

### Strategy
Create new array, place each element at rotated position.

### Algorithm
```
1. Create temp array of size n
2. For i from 0 to n-1:
   - temp[(i + k) % n] = nums[i]
3. Copy temp back to nums
```

### Complexity
- **Time:** O(n)
- **Space:** O(n) ❌ (doesn't meet follow-up requirement)

## ⏱️ Complexity Analysis

### **Approach 1: Three Reversals (Optimal)**
**Time Complexity: O(n)**
- Reverse entire array: O(n)
- Reverse first k: O(k)
- Reverse last n-k: O(n-k)
- Total: O(n) + O(k) + O(n-k) = O(2n) = O(n)

**Space Complexity: O(1)**
- Only swap variables
- In-place modification
- Meets follow-up requirement ✓

### **Comparison:**
| Approach | Time | Space | In-place? |
|----------|------|-------|-----------|
| Three Reversals | O(n) | O(1) | ✓ |
| Cyclic Replace | O(n) | O(1) | ✓ |
| Extra Array | O(n) | O(n) | ✗ |

## ⚠️ Common Pitfalls

1. **Forgetting k % n**
   - Mistake: k = 10, n = 7 → rotate 10 times
   - Reality: 10 % 7 = 3, only need to rotate 3 times
   - Example: Rotating by n is same as no rotation

2. **Off-by-one in reverse ranges**
   - Mistake: `reverse(0, k)` instead of `reverse(0, k-1)`
   - Issue: Includes element at index k, wrong range
   - Fix: Use inclusive ranges carefully

3. **Wrong reversal order**
   - Mistake: Reverse parts first, then entire array
   - Issue: Doesn't produce correct result
   - Fix: Must reverse entire array FIRST

4. **Not handling k = 0 or k = n**
   - Edge case: k=0 → no rotation needed
   - Edge case: k=n → full rotation = original
   - Fix: k % n handles both automatically

## 🔑 Pattern Recognition

**Pattern: Array Reversal**
- Swap elements from both ends moving toward center
- O(1) space, O(n) time
- Useful for in-place transformations

**When to Use Reversal Pattern:**
- ✅ Need to reverse order of elements
- ✅ Need to swap chunks of array
- ✅ In-place requirement with O(1) space
- ✅ "Rotate" or "shift" operations

**Similar Problems:**
- Reverse String (#344)
- Reverse Words in a String (#151)
- Rotate String (#796)
- Reverse Linked List (#206) - Same concept, different structure

**Problem Variations:**
- Rotate left instead of right → Use k = n - k
- Rotate 2D matrix → Apply same reversal logic per row/column
- Rotate linked list → Need different approach (no random access)

**Recognition Triggers:**
- See "rotate array" → Think three reversals
- See "in-place with O(1) space" → Reversal likely optimal
- See "move chunks" → Consider reversal approach

## 🎓 Alternative Thinking

**Why not just move elements?**
```
Naive approach: 
for i from 0 to k:
    temp = nums[n-1]
    shift all elements right
    nums[0] = temp
```
**Problem:** O(n*k) time - too slow!

**Better:** Reversal is O(n) regardless of k size.

**Key insight:** Don't think about k rotations, think about final positions!
