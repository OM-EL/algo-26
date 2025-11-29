# Two Sum Visual Schema

## Problem Visualization

```
Array:    ┌───┬───┬────┬────┐
  nums =  │ 2 │ 7 │ 11 │ 15 │
          └───┴───┴────┴────┘
Index:      0   1    2    3

Target: 9

Question: Which two numbers sum to 9?
```

## Brute Force Approach (O(n²))

```
Check ALL pairs:

nums[0] + nums[1] = 2 + 7  = 9  ✓ FOUND!
nums[0] + nums[2] = 2 + 11 = 13 ✗
nums[0] + nums[3] = 2 + 15 = 17 ✗
nums[1] + nums[2] = 7 + 11 = 18 ✗
nums[1] + nums[3] = 7 + 15 = 22 ✗
nums[2] + nums[3] = 11 + 15= 26 ✗

Total comparisons: n(n-1)/2 = 6 for n=4
```

## Hash Map Approach (O(n))

```
Step-by-step execution:

┌─────────────────────────────────────────────────────────────┐
│ Step 1: num = 2 (index 0)                                   │
│                                                             │
│   complement = 9 - 2 = 7                                    │
│   Is 7 in map? NO                                           │
│   Add to map: {2: 0}                                        │
│                                                             │
│   Map State: ┌─────┬───────┐                                │
│              │ Key │ Index │                                │
│              ├─────┼───────┤                                │
│              │  2  │   0   │                                │
│              └─────┴───────┘                                │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ Step 2: num = 7 (index 1)                                   │
│                                                             │
│   complement = 9 - 7 = 2                                    │
│   Is 2 in map? YES! at index 0                              │
│                                                             │
│   RETURN [0, 1] ✓                                           │
└─────────────────────────────────────────────────────────────┘
```

## Complexity Comparison

```
                    Time        Space
                    ─────────   ─────────
Brute Force:        O(n²)       O(1)
                    │           │
                    │           └── No extra space
                    └── Check all pairs

Hash Map:           O(n)        O(n)
                    │           │
                    │           └── Store up to n elements
                    └── Single pass, O(1) lookup
```

## Memory Layout (Hash Map)

```
Hash Map Internal Structure:

┌──────────────────────────────────────┐
│          Hash Function               │
│                                      │
│    key ──────► hash(key) % size      │
└──────────────────────────────────────┘
                    │
                    ▼
┌─────┬─────┬─────┬─────┬─────┬─────┐
│  0  │  1  │  2  │  3  │  4  │ ... │  Buckets
├─────┼─────┼─────┼─────┼─────┼─────┤
│     │     │ 2→0 │     │ 7→1 │     │
│     │     │     │     │     │     │
└─────┴─────┴─────┴─────┴─────┴─────┘

Lookup: hash(7) = 4 → bucket[4] → found {7: 1}
```

## Decision Tree

```
Start
  │
  ▼
┌─────────────────────┐
│ Is array sorted?    │
└─────────────────────┘
     │          │
    Yes        No
     │          │
     ▼          ▼
┌─────────┐  ┌─────────────┐
│Two      │  │ Use Hash    │
│Pointers │  │ Map O(n)    │
│O(n)     │  │             │
└─────────┘  └─────────────┘
```
