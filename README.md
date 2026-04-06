# DSA in Go 🧠

A collection of fundamental Data Structures and Algorithms implemented in Go, organized by topic.

---

## 📁 Project Structure

```
DSA/
├── main.go
├── binary-tree/
│   └── binary_traversals.go
├── binary-searchT/
│   └── bst_insert.go
├── heap/
│   ├── heap_creation.go
│   ├── heap_container.go
│   └── heap_priorityQ.go
├── hash-map/
│   └── hash.go
└── sorting/
    ├── insert_sort.go
    ├── selection_sort.go
    └── buble_sort.go
```

---

## 📦 Packages

### 🌳 `binary-tree`
Full binary tree implementation with the following operations:

| Function | Description |
|---|---|
| `BuildTree(preOrder, inOrder)` | Constructs a binary tree from preorder and inorder traversals |
| `InOrder(root)` | Left → Root → Right traversal |
| `PreOrder(root)` | Root → Left → Right traversal |
| `LevelPrint(root)` | BFS level-order traversal (flat) |
| `LevelPrintPyramid(root)` | BFS level-order traversal (level by level) |
| `Insert(root, val)` | Inserts a node at the first available position (BFS) |
| `Search(root, k)` | DFS search for a value |
| `Searchh(root, val)` | BFS search for a value |
| `Update(root, oldVal, newVal)` | Updates a node's value (BFS) |
| `Delete(root, value)` | Deletes a node and replaces it with the deepest rightmost node |
| `Height(root)` | Returns the height of the tree |
| `CountNode(root)` | Returns total number of nodes |
| `CountLeaf(root)` | Returns number of leaf nodes |
| `CountInterval(root)` | Returns number of internal nodes |
| `Diameter(root)` | Returns the diameter (longest path between any two nodes) |

---

### 🔍 `binary-searchT`
Binary Search Tree (BST) implementation.

| Function | Description |
|---|---|
| `Insert(root, val)` | Inserts a value maintaining BST property |
| `Inorder(root)` | Prints nodes in sorted (ascending) order |

> **Note:** `BuildTree(preOrder, inOrder)` from the binary-tree package requires a **sorted inorder** array — which is always the case for a valid BST.

---

### 🏔️ `heap`

Three heap implementations:

#### `MinHeap` (manual implementation)
A min-heap built from scratch without using `container/heap`.

| Method | Description |
|---|---|
| `Insert(v)` | Inserts a value and heapifies up |
| `Extract()` | Removes and returns the minimum element |
| `HeapifyUp(index)` | Bubbles element up to maintain heap property |
| `HeapifyDown(index)` | Bubbles element down to maintain heap property |
| `Display()` | Prints all heap elements |

#### `InHeap` (implements `container/heap` interface)
A min-heap using Go's standard `container/heap` package.

```go
h := &myheap.InHeap{4, 1, 2, 8, 7}
heap.Init(h)
heap.Push(h, 10)
min := (*h)[0]  // peek at min
val := heap.Pop(h).(int)
```

#### `Pq` — Priority Queue (Max Heap)
A max-heap priority queue using `container/heap`.

```go
p := &myheap.Pq{
    {Val: "Task1", Priority: 1},
    {Val: "Task2", Priority: 5},
}
heap.Init(p)
heap.Push(p, &myheap.MaxHeap{Val: "Task3", Priority: 10})
item := heap.Pop(p).(*myheap.MaxHeap)
```

> Items with **higher priority values** are popped first.

---

### 🗃️ `hash-map`
A hash map implementation using **separate chaining** for collision resolution.

| Function / Method | Description |
|---|---|
| `BuildHashMap(size)` | Creates a new hash map with the given number of buckets |
| `Put(key, value)` | Inserts or updates a key-value pair |
| `Get(key)` | Returns the value and a boolean indicating if key exists |
| `Remove(key)` | Deletes a key-value pair |
| `Display()` | Prints all buckets and their entries |
| `HashFunc(key)` | Computes bucket index by summing character ASCII values mod size |

---

### 🔢 `sorting`
Sorting algorithms on integer slices (ascending order).

| Function | Algorithm | Time Complexity |
|---|---|---|
| `Insertion(arr)` | Insertion Sort | O(n²) avg, O(n) best |
| `Selection(arr)` | Selection Sort | O(n²) |

---

## 🚀 Running the Project

```bash
# From the project root
go run main.go
```

Make sure your `go.mod` module name matches the import paths (e.g., `module dsa`).

```
# go.mod example
module dsa

go 1.21
```

---

## 📝 Notes

- The `buble_sort.go` file exists in the sorting package but does not yet contain an implementation.
- The `main.go` in `heap/heap_creation.go` contains a standalone `main()` for isolated testing — move it to a `_test.go` or remove it to avoid build conflicts.
- All heap operations using `container/heap` require type assertions when popping: `heap.Pop(h).(int)` or `heap.Pop(p).(*MaxHeap)`.
