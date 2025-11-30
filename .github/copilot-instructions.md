# GitHub Copilot Instructions for algo-26

## 📋 Project Overview

This is an algorithm learning repository organized by problem categories. Each exercise follows a **strict 3-file structure** for clarity and focused learning.

## 🏗️ Repository Structure

```
algo-26/
├── exercises/
│   ├── arrays/
│   │   ├── two_sum/
│   │   │   ├── README.md       # Problem statement (énoncé)
│   │   │   ├── NOTES.md        # Solutions + Decision tree
│   │   │   └── LITERATURE.md   # Historical context (short & dense)
│   │   └── remove_duplicates_sorted_array/
│   │       ├── README.md
│   │       ├── NOTES.md
│   │       └── LITERATURE.md
│   ├── backtracking/
│   ├── dynamic_programming/
│   ├── graphs/
│   ├── trees/
│   └── ... (other categories)
├── templates/
│   └── exercise_template/
└── docs/
```

## 📁 Exercise Structure (STRICT FORMAT)

Each exercise **MUST** contain exactly **3 files**:

### 1. **README.md** (Énoncé / Problem Statement)
- Problem description
- Examples (2-3 with input/output)
- Constraints
- Follow-up questions
- **Difficulty & Category tags**
- **Source** (e.g., LeetCode #26)

**NO CODE - Only problem description!**

### 2. **NOTES.md** (Solution Explanations + Decision Tree)

**Purpose:** Teach HOW to think about the problem, not just the solution.

**Structure (in order):**

1. 🎯 **Core Insight** (1-2 sentences)
   - The "aha!" moment that unlocks the problem
   - What property makes the solution possible?

2. 🔍 **Decision Tree** (ASCII diagram)
   - Shows decision-making process
   - "If X, then approach Y; if Z, then approach W"
   - Helps understand WHEN to use this pattern

3. 💭 **How to Think About This**
   - First observations when reading the problem
   - Questions to ask yourself
   - What constraints matter most?
   - What makes this problem similar/different from others?

4. 💡 **Solution Approach** (Step-by-step)
   - Explain the strategy in plain language first
   - Why this approach works (intuition)
   - Algorithm steps (numbered, clear)
   - Visual example with concrete numbers

5. 🧮 **Why This Works** (Mathematical/Logical Proof)
   - Brief explanation of correctness
   - Why greedy/DP/two-pointers is optimal here
   - Connect back to core insight

6. ⏱️ **Complexity Analysis**
   - Time complexity with explanation
   - Space complexity with explanation
   - Why this is optimal (or trade-offs made)

7. ⚠️ **Common Pitfalls** (3-4 mistakes)
   - What beginners typically get wrong
   - Off-by-one errors
   - Edge cases often missed

8. 🔑 **Pattern Recognition**
   - What pattern/technique is this?
   - When to apply this pattern?
   - Similar problems that use same approach

**Writing Style:**
- **Clear over clever** - explain like teaching a friend
- **Concise but complete** - no fluff, but don't skip steps
- **Visual aids** - ASCII diagrams, worked examples
- **Logical flow** - each section builds on previous
- **NO code** - focus on understanding, not implementation

### 3. **LITERATURE.md** (Historical Context)
- 📚 **Historical Background** (when/why technique was developed)
- 🎓 **Theoretical Foundations** (complexity theory, proofs)
- 🔗 **Related Problems & Patterns** (connections to other algorithms)
- 🌍 **Real-World Applications** (databases, compression, etc.)
- 💻 **Performance Insights** (cache, CPU behavior)
- 📖 **References** (books, papers)

**Style:**
- **Short & dense** (not overwhelming)
- Focus on context, not repetition
- 80-100 lines maximum
- Academic but accessible

## 🚫 What NOT to Include

- ❌ **NO code files** (`solution.go`, `solution.py`, etc.)
- ❌ **NO test files** (`solution_test.go`, etc.)
- ❌ **NO schemas directory** (keep visuals in NOTES.md)
- ❌ **NO verbose explanations** (be concise)
- ❌ **NO redundant content** across files

## ✅ When Creating New Exercises

1. **Identify the category** (arrays, trees, graphs, etc.)
2. **Create directory** under `exercises/<category>/<problem_name>/`
3. **Generate 3 files only**: README.md, NOTES.md, LITERATURE.md
4. **Follow the exact structure** shown above
5. **Keep it concise** - quality over quantity

## 📝 Naming Conventions

- **Directories:** `snake_case` (e.g., `two_sum`, `remove_duplicates_sorted_array`)
- **Files:** Always `README.md`, `NOTES.md`, `LITERATURE.md` (uppercase)
- **Categories:** Lowercase, underscored (e.g., `dynamic_programming`, `bit_manipulation`)

## 🎯 Design Philosophy

- **Clarity over completeness** - Focus on understanding, not exhaustive coverage
- **Visual over verbal** - Use ASCII diagrams, not long paragraphs
- **Dense over verbose** - Every sentence should add value
- **Structured over free-form** - Consistent format aids learning

## 🔧 Category Organization

Current categories:
- `arrays` - Array manipulation problems
- `backtracking` - Recursive exploration
- `bit_manipulation` - Bitwise operations
- `dynamic_programming` - Memoization & tabulation
- `graphs` - Graph traversal & algorithms
- `greedy` - Greedy choice strategies
- `hash_table` - Hashing problems
- `heap` - Priority queue problems
- `linked_lists` - Linked list manipulation
- `math` - Mathematical algorithms
- `queue` - Queue-based problems
- `recursion` - Recursive solutions
- `searching` - Search algorithms
- `sliding_window` - Window-based problems
- `sorting` - Sorting algorithms
- `stack` - Stack-based problems
- `strings` - String manipulation
- `trees` - Tree traversal & algorithms
- `two_pointers` - Two-pointer technique

## 💡 Example: Adding a New Problem

If user says: "Add problem X"

1. **Ask clarifying questions** (if needed):
   - Category?
   - Difficulty?
   - Source (LeetCode #)?

2. **Create 3 files** in appropriate category directory

3. **README.md**: Problem statement from source

4. **NOTES.md**: Follow this thought process:
   - What's the key insight? (Start here!)
   - What questions should I ask when seeing this?
   - What approach does the insight suggest?
   - How do I decide between approaches?
   - Walk through a concrete example
   - Why does this work mathematically?
   - What mistakes do people make?
   - What pattern is this, and when do I use it again?

5. **LITERATURE.md**: Historical context + applications

6. **No code, no tests, no extra files!**

## 🧠 Teaching Philosophy for NOTES.md

**Goal:** Enable independent problem-solving, not memorization.

**Good NOTES.md answers:**
- ✅ "When I see X constraint, I should think Y"
- ✅ "This is similar to Z problem because..."
- ✅ "The insight is that [property] means we can..."
- ✅ "Ask yourself: Is the input sorted? Can I use that?"

**Bad NOTES.md:**
- ❌ Just listing algorithm steps without explanation
- ❌ Jumping to solution without explaining thought process
- ❌ Missing the "why" behind the approach
- ❌ Not connecting to broader patterns

**Framework:**
```
See problem → Ask questions → Notice patterns → 
Choose approach → Understand why → Recognize for future
```

## 🎓 Learning Objectives

This repository prioritizes:
- **Understanding patterns** over memorizing solutions
- **Decision-making skills** (when to use which approach)
- **Historical context** (why algorithms were developed)
- **Real-world connections** (where algorithms are applied)

## 🔄 Maintenance

- Keep files **up to date** with best practices
- **Refactor** when structure improves
- **Remove redundancy** across files
- Ensure **consistency** in formatting

---

**Remember:** Less is more. Focus on clarity, structure, and insight.
