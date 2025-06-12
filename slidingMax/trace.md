# Sliding Window Maximum - Optimized Solution Trace

### Input:
`nums = [1, 3, -1, -3, 5, 3, 6, 7], k = 3`

### Step-by-Step Execution:

#### **Iteration 1** (`i = 0`):
- Current element: `nums[0] = 1`
- **Deque Operations**:
  - Add index `0` to the deque.
- **Deque Visualization**:
```
Deque: [0]
Result: []
```

---

#### **Iteration 2** (`i = 1`):
- Current element: `nums[1] = 3`
- **Deque Operations**:
  - Remove index `0` (since `nums[0] < nums[1]`).
  - Add index `1` to the deque.
- **Deque Visualization**:
```
Deque: [1]
Result: []
```

---

#### **Iteration 3** (`i = 2`):
- Current element: `nums[2] = -1`
- **Deque Operations**:
  - Add index `2` to the deque.
- **Deque Visualization**:
```
Deque: [1, 2]
Result: [3]
```

---

#### **Iteration 4** (`i = 3`):
- Current element: `nums[3] = -3`
- **Deque Operations**:
  - Remove index `1` (out of bounds for the current window).
  - Add index `3` to the deque.
- **Deque Visualization**:
```
Deque: [2, 3]
Result: [3, 3]
```

---

#### **Iteration 5** (`i = 4`):
- Current element: `nums[4] = 5`
- **Deque Operations**:
  - Remove indices `2` and `3` (since `nums[2] < nums[4]` and `nums[3] < nums[4]`).
  - Add index `4` to the deque.
- **Deque Visualization**:
```
Deque: [4]
Result: [3, 3, 5]
```

---

#### **Iteration 6** (`i = 5`):
- Current element: `nums[5] = 3`
- **Deque Operations**:
  - Add index `5` to the deque.
- **Deque Visualization**:
```
Deque: [4, 5]
Result: [3, 3, 5, 5]
```

---

#### **Iteration 7** (`i = 6`):
- Current element: `nums[6] = 6`
- **Deque Operations**:
  - Remove indices `4` and `5` (since `nums[4] < nums[6]` and `nums[5] < nums[6]`).
  - Add index `6` to the deque.
- **Deque Visualization**:
```
Deque: [6]
Result: [3, 3, 5, 5, 6]
```

---

#### **Iteration 8** (`i = 7`):
- Current element: `nums[7] = 7`
- **Deque Operations**:
  - Remove index `6` (since `nums[6] < nums[7]`).
  - Add index `7` to the deque.
- **Deque Visualization**:
```
Deque: [7]
Result: [3, 3, 5, 5, 6, 7]
```

---

### Final Result:
`result = [3, 3, 5, 5, 6, 7]`

This step-by-step breakdown includes the state of the result array alongside the deque visualization during each iteration.
