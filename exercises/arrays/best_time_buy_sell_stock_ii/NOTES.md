# Solution Explanations & Decision Tree

## 🎯 Core Insight

**Capture every upward price movement!**
```
If prices[i+1] > prices[i] → Add (prices[i+1] - prices[i]) to profit
```

**Why this works:** Multiple small transactions = one big transaction when prices only go up.

## 🔍 Decision Tree

```
Start: Maximize profit with unlimited transactions
                │
                ▼
        ┌───────────────┐
        │ Can we hold    │
        │ multiple       │
        │ stocks?        │
        └───────────────┘
             │      │
            No     Yes → Different problem
             │
             ▼
        ┌───────────────┐
        │ Unlimited      │
        │ transactions?  │
        └───────────────┘
             │      │
           Yes     No → Stock I, III, IV (DP needed)
             │
             ▼
        ┌──────────────────┐
        │ Greedy Approach  │
        │ Capture all      │
        │ upward moves     │
        │ Time: O(n)       │
        │ Space: O(1)      │
        └──────────────────┘
```

## 💭 How to Think About This

**First Observations:**
1. "Unlimited transactions" - no limit on buy/sell pairs
2. "At most one share" - must sell before buying again
3. Goal: maximize total profit

**Questions to Ask Yourself:**
- Q: What if I could make every profitable move?
- A: That would be optimal since no transaction limit!

- Q: Does timing of buy/sell matter if I can transact unlimited times?
- A: No! Only price differences matter.

- Q: Can I break down a long hold into smaller transactions?
- A: Yes! Buy at 1, sell at 5 = (5-1) = (2-1)+(3-2)+(4-3)+(5-4)

**Key Realization:**
```
Holding from day i to day j:
profit = price[j] - price[i]

Same as multiple transactions:
profit = (price[i+1] - price[i]) + (price[i+2] - price[i+1]) + ... + (price[j] - price[j-1])
```

These are mathematically equivalent! So just sum all positive differences.

## 💡 Solution Approach: Greedy

### Strategy (Plain Language)
Walk through prices day by day. Whenever tomorrow's price is higher than today's, imagine buying today and selling tomorrow. Sum all these gains.

### Why Greedy Works Here
- No transaction limit → can split any transaction into smaller ones
- No cost per transaction → all gains worth taking
- Can't predict future → but don't need to! Every gain is independent

**Intuition:**
```
Prices: [1, 3, 2, 4]

Optimal thinking:
"3 > 1? Take it! (+2)"
"2 < 3? Skip."
"4 > 2? Take it! (+2)"
Total: 4

This IS the optimal solution!
```

### Algorithm
```
1. profit = 0
2. For each day i from 1 to n-1:
   - If prices[i] > prices[i-1]:
     → profit += (prices[i] - prices[i-1])
3. Return profit
```

### Visual Example
```
prices = [7, 1, 5, 3, 6, 4]

Day 0→1: 1 < 7   → Skip (would lose money)
Day 1→2: 5 > 1   → +4 (BUY at 1, SELL at 5)
Day 2→3: 3 < 5   → Skip (would lose money)
Day 3→4: 6 > 3   → +3 (BUY at 3, SELL at 6)
Day 4→5: 4 < 6   → Skip (would lose money)

Total profit: 7 ✓

Graph visualization:
  7 •
    │ ╲
  6 │   ╲     • ← Sell here
  5 │     • ← Sell here
  4 │     │ ╲ •
  3 │     │   • ← Buy here
  2 │     │
  1 │     • ← Buy here
  0 └─────────────────────
    0  1  2  3  4  5  (days)

Notice: We capture the two "hills" (1→5 and 3→6)
```

## 🧮 Why This Works (Mathematical Proof)

**Claim:** Summing all positive differences gives optimal profit.

**Proof:**
1. Any valid transaction sequence can be represented as a sum of daily differences
2. For a buy at day i and sell at day j:
   ```
   profit = price[j] - price[i]
          = (price[j] - price[j-1]) + (price[j-1] - price[j-2]) + ... + (price[i+1] - price[i])
          = Σ(price[k+1] - price[k]) for k ∈ [i, j-1]
   ```

3. Taking only positive differences means:
   - Include all upward moves (gains)
   - Exclude all downward moves (losses)

4. This is optimal because:
   - Including a negative difference would reduce profit
   - Missing a positive difference would miss potential gain
   - With unlimited transactions, we can take all gains independently

**Greedy Choice Property:** Each local decision (take this gain or not) is optimal globally.

## ⏱️ Complexity Analysis

**Time Complexity: O(n)**
- Single pass through prices array
- One comparison per adjacent pair
- n-1 comparisons total

**Space Complexity: O(1)**
- Only one variable (profit) needed
- No arrays, no recursion stack
- Input array not modified

**Optimality:** Can't do better than O(n) since we must examine every price at least once.

## ⚠️ Common Pitfalls

1. **Overthinking with Dynamic Programming**
   - Mistake: "This looks like a DP problem..."
   - Reality: Unlimited transactions make it greedy
   - Stock I, III, IV need DP because of transaction limits

2. **Trying to find global peaks/valleys**
   - Mistake: "I need to find THE lowest buy and highest sell"
   - Reality: Multiple small transactions can be better
   - Example: [1,2,3,4,5] → buying at 1, selling daily is optimal

3. **Forgetting zero profit is valid**
   - Mistake: Always trying to make a transaction
   - Reality: If prices always decrease, don't trade (profit = 0)
   - Example: [5,4,3,2,1] → correct answer is 0

4. **Off-by-one errors in loop**
   - Mistake: `for i from 0 to n-1` comparing `prices[i]` with `prices[i+1]`
   - Issue: `prices[i+1]` goes out of bounds
   - Fix: `for i from 1 to n-1` comparing `prices[i]` with `prices[i-1]`

## 🔑 Pattern Recognition

**Pattern: Greedy Algorithm**
- Make locally optimal choice at each step
- Don't need to reconsider previous choices
- Local optima lead to global optimum

**When to Use Greedy:**
- ✅ No dependency between choices (can take gains independently)
- ✅ Optimal substructure (subproblem solutions combine optimally)
- ✅ Greedy choice property (local optimal → global optimal)

**When NOT to Use Greedy:**
- ❌ Limited transactions (need DP to track states)
- ❌ Additional constraints (cooldown, fees - need DP)
- ❌ Choices affect future options (need to explore states)

**Similar Greedy Problems:**
- Jump Game II (#45) - minimize jumps
- Gas Station (#134) - find starting point
- Partition Labels (#763) - partition by character ranges

**Stock Series (Know the Difference!):**
| Problem | Transactions | Approach | Complexity |
|---------|-------------|----------|------------|
| Stock I (#121) | 1 | Track min, find max diff | O(n), O(1) |
| **Stock II (#122)** | **Unlimited** | **Greedy (sum gains)** | **O(n), O(1)** |
| Stock III (#123) | 2 | DP with states | O(n), O(1) |
| Stock IV (#188) | k | DP with k states | O(nk), O(k) |
| Cooldown (#309) | Unlimited + cooldown | State machine DP | O(n), O(1) |
| Fee (#714) | Unlimited + fee | Modified greedy/DP | O(n), O(1) |

**Recognition Triggers:**
- See "unlimited transactions" → Think greedy, sum all gains
- See "at most k transactions" → Think DP with state tracking
- See "additional constraints" → Think DP with states
