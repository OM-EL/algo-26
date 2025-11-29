# Algo-26: Coding Exercises & Problem-Solving Library

A comprehensive collection of coding exercises (LeetCode-style) implemented in Go, with detailed explanations, schemas, problem-solving strategies, and historical context.

## 📁 Project Structure

```
algo-26/
├── exercises/                    # All coding exercises organized by category
│   ├── arrays/                   # Array manipulation problems
│   ├── strings/                  # String manipulation problems
│   ├── linked_lists/             # Linked list problems
│   ├── trees/                    # Binary trees, BST, etc.
│   ├── graphs/                   # Graph traversal and algorithms
│   ├── dynamic_programming/      # DP problems
│   ├── sorting/                  # Sorting algorithms
│   ├── searching/                # Search algorithms
│   ├── two_pointers/             # Two pointer technique
│   ├── sliding_window/           # Sliding window technique
│   ├── stack/                    # Stack-based problems
│   ├── queue/                    # Queue-based problems
│   ├── heap/                     # Heap/Priority queue problems
│   ├── hash_table/               # Hash table problems
│   ├── recursion/                # Recursive problems
│   ├── backtracking/             # Backtracking problems
│   ├── greedy/                   # Greedy algorithm problems
│   ├── bit_manipulation/         # Bit manipulation problems
│   └── math/                     # Mathematical problems
├── docs/
│   └── schemas/                  # Visual diagrams and schemas
├── templates/                    # Templates for new exercises
├── go.mod                        # Go module file
└── README.md                     # This file
```

## 📝 Exercise Structure

Each exercise follows a standardized structure:

```
exercises/<category>/<problem_name>/
├── README.md              # Problem description, examples, constraints
├── solution.go            # Go implementation
├── solution_test.go       # Unit tests
├── NOTES.md               # Problem-solving approach and thinking process
├── LITERATURE.md          # Historical context and related algorithms
└── schemas/               # Visual diagrams specific to this problem
    └── approach.png
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests for a specific category
go test ./exercises/arrays/...

# Run tests with verbose output
go test -v ./exercises/arrays/two_sum/...
```

### Adding a New Exercise

1. Copy the template from `templates/exercise_template/`
2. Place it in the appropriate category directory
3. Fill in the problem description, solution, notes, and literature

## 📚 Categories Overview

| Category | Description | Difficulty Range |
|----------|-------------|------------------|
| Arrays | Array traversal, manipulation, subarrays | Easy - Hard |
| Strings | String manipulation, pattern matching | Easy - Hard |
| Linked Lists | Singly/doubly linked lists, operations | Easy - Medium |
| Trees | Binary trees, BST, tree traversals | Easy - Hard |
| Graphs | DFS, BFS, shortest path, topological sort | Medium - Hard |
| Dynamic Programming | Memoization, tabulation, state transitions | Medium - Hard |
| Sorting | Various sorting algorithms and applications | Easy - Medium |
| Searching | Binary search and variations | Easy - Medium |
| Two Pointers | Two pointer technique problems | Easy - Medium |
| Sliding Window | Fixed/variable window problems | Medium - Hard |
| Stack | Stack operations, monotonic stack | Easy - Hard |
| Queue | Queue operations, BFS applications | Easy - Medium |
| Heap | Priority queues, top K problems | Medium - Hard |
| Hash Table | Hashing, frequency counting | Easy - Medium |
| Recursion | Recursive problem solving | Easy - Hard |
| Backtracking | Permutations, combinations, subsets | Medium - Hard |
| Greedy | Greedy algorithm applications | Medium - Hard |
| Bit Manipulation | Bitwise operations | Easy - Hard |
| Math | Mathematical algorithms | Easy - Hard |

## 📖 Documentation Standards

### Problem Description (README.md)
- Clear problem statement
- Input/Output examples
- Constraints
- Follow-up questions (if any)

### Solution Notes (NOTES.md)
- Initial thought process
- Different approaches considered
- Time/Space complexity analysis
- Edge cases
- Key insights

### Literature (LITERATURE.md)
- Historical context
- Related algorithms and data structures
- Academic references
- Real-world applications

## 🤝 Contributing

1. Follow the exercise template structure
2. Include comprehensive test cases
3. Document your problem-solving approach
4. Add relevant literature and historical context

## 📜 License

MIT License
