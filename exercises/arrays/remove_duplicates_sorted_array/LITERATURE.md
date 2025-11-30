# Literature & Context

## 📚 Historical Background

**In-Place Algorithms (1959-1970s)**
- Memory was scarce and expensive (measured in KB)
- Tony Hoare's QuickSort (1959) pioneered in-place sorting
- Space complexity became as critical as time complexity

**Two-Pointer Technique (1960s-1980s)**
- Emerged from Merge Sort's merge process (1962)
- Formalized as a distinct pattern in the 1980s
- Key insight: Multiple pointers with different roles (read/write, start/end)

## 🎓 Theoretical Foundations

**Loop Invariant (Tony Hoare)**
- At each iteration: `nums[0..slow]` contains unique elements
- Proves correctness formally
- Initialization → Maintenance → Termination

**Complexity Bounds**
- **Time:** Θ(n) - must examine every element
- **Space:** O(1) auxiliary - only pointers, no new structures
- **Amortized:** 2n operations (read once, write at most once)

## 🔗 Related Problems & Patterns

**Similar Two-Pointer Problems:**
- Remove Element (#27) - Remove specific value
- Move Zeroes (#283) - Partition array
- Remove Duplicates II (#80) - Allow k duplicates
- Partition (Quicksort) - Lomuto scheme uses slow/fast

**Sorted Array Properties:**
- Monotonicity: Each element ≥ previous
- Duplicate clustering: All duplicates adjacent
- Enables O(n) deduplication vs O(n²) or O(n) space unsorted

## 🌍 Real-World Applications

**Databases**
- `SELECT DISTINCT` implementation (sorted results)
- Index deduplication after bulk inserts
- Log compaction in write-ahead logs

**Data Compression**
- Run-Length Encoding (RLE): Remove consecutive duplicates
- Used in BMP, PCX, fax transmission (T.4/T.6)
- Video compression early codecs

**File Systems**
- ZFS, Btrfs block deduplication
- Process: Hash blocks → Sort → Deduplicate → Save 50-90% storage

**Stream Processing**
- Log aggregation deduplication
- Time-series sensor data cleanup
- Network packet duplicate removal

## 💻 Performance Insights

**CPU Cache Efficiency**
- Sequential access = 1-2 cycles/element (cache hits)
- Prefetcher loads nearby elements automatically
- Random access = 200-300 cycles (cache misses)

**Branch Prediction**
- Modern CPUs predict `if nums[fast] != nums[slow]`
- Many duplicates → predictable not-taken
- No duplicates → predictable taken

## 📖 References

- **Knuth, "TAOCP Vol 3"** - In-place algorithms
- **CLRS** - Loop invariants & correctness proofs
- **Bentley, "Programming Pearls"** - Practical two-pointer techniques
