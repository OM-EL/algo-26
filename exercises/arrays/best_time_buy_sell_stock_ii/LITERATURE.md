# Literature & Context

## 📚 Historical Background

**Greedy Algorithms (1950s-1970s)**
- Formalized by various researchers in operations research
- Kruskal's algorithm (1956) and Prim's algorithm (1957) for MST
- Dijkstra's shortest path (1959) - greedy paradigm
- Huffman coding (1952) - optimal greedy solution

**Stock Trading Algorithms**
- Academic interest since 1960s with emergence of algorithmic trading
- Practical importance grew with electronic markets (1970s-1980s)
- High-frequency trading made O(n) algorithms essential

## 🎓 Theoretical Foundations

**Greedy Choice Property**
- A global optimum can be reached by making locally optimal choices
- Not all problems have this property (requires proof)

**Why Greedy Works Here:**
```
Proof sketch:
- Any sequence of buy/sell transactions can be decomposed
- (sell_j - buy_i) = Σ(price[k+1] - price[k]) for k in [i, j-1]
- Summing all positive differences = optimal
- Missing any positive diff = suboptimal
```

**Matroid Theory**
- Greedy algorithms work on problems with matroid structure
- Stock problem satisfies exchange property
- Locally optimal → globally optimal

**Complexity Class**
- P (polynomial time) - efficiently solvable
- Contrast with constrained versions (k transactions) requiring DP

## 🔗 Related Problems & Patterns

**Stock Series Progression:**
1. **Stock I** (#121) - One transaction → O(n) single pass
2. **Stock II** (#122) - Unlimited → Greedy O(n)
3. **Stock III** (#123) - Two transactions → DP O(n)
4. **Stock IV** (#188) - k transactions → DP O(nk)
5. **With Cooldown** (#309) - State machine DP
6. **With Fee** (#714) - Modified greedy/DP

**Greedy Paradigm:**
- Activity selection problem
- Fractional knapsack
- Huffman encoding
- Minimum spanning tree (Kruskal, Prim)

**Peak-Valley Pattern:**
- Also used in signal processing
- Time series analysis
- Technical analysis in finance

## 🌍 Real-World Applications

**Algorithmic Trading**
- High-frequency trading strategies
- Market-making algorithms
- Momentum trading (buy on uptrend)
- Real systems use sophisticated versions with fees, slippage

**Financial Analysis**
- Backtesting trading strategies
- Performance benchmarking
- Risk analysis (maximum drawdown)
- Portfolio optimization

**Technical Indicators**
- Moving Average Convergence Divergence (MACD)
- Relative Strength Index (RSI)
- Both analyze price momentum like our algorithm

**Resource Allocation**
- Server capacity planning (buy when cheap)
- Energy trading (store when cheap, sell when expensive)
- Inventory management

## 💻 Performance Insights

**Single Pass Efficiency**
- O(n) = optimal (must examine all prices)
- O(1) space = cache-friendly
- Branch prediction: if statement is data-dependent

**Real Trading Constraints**
- Transaction fees → not all gains profitable
- Slippage → execution price differs from observed price
- Market impact → large orders move prices
- Latency → delays between decision and execution

**Practical Modifications**
```
profit += max(0, prices[i] - prices[i-1] - fee)
```
Add transaction fee consideration.

## 📖 References

- **Cormen (CLRS)** - Greedy algorithms chapter, matroid theory
- **"Introduction to Algorithms"** - Proof techniques for greedy correctness
- **Dasgupta, Papadimitriou, Vazirani** - "Algorithms" chapter on greedy methods
- **Finance Literature** - "A Random Walk Down Wall Street" by Malkiel
