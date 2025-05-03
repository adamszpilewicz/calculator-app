package calculator

import (
	"container/heap"
	"sort"
)

// PackResult represents a pack size and the count of that pack
type PackResult struct {
	Size  int `json:"size"`
	Count int `json:"count"`
}

// CalculatePacks calculates the optimal pack combination
func CalculatePacks(quantity int, packSizes []int) []PackResult {
	// Priority queue for Dijkstra-style search
	pq := &priorityQueue{}
	heap.Init(pq)

	// Track visited totals to avoid revisiting same total
	visited := make(map[int]bool)

	// Start from total = 0
	heap.Push(pq, &state{
		totalItems: 0,
		numPacks:   0,
		packCount:  make(map[int]int),
	})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*state)

		// If current total fulfills or exceeds quantity, we have a valid solution
		if curr.totalItems >= quantity {
			// Convert map to []PackResult
			var results []PackResult
			for size, count := range curr.packCount {
				results = append(results, PackResult{Size: size, Count: count})
			}

			// Sort results by descending size for consistent output
			sort.Slice(results, func(i, j int) bool {
				return results[i].Size > results[j].Size
			})

			return results
		}

		// Explore next states by adding one more pack of each size
		for _, size := range packSizes {
			nextTotal := curr.totalItems + size

			if visited[nextTotal] {
				continue // already visited this total
			}
			visited[nextTotal] = true

			newPackCount := copyMap(curr.packCount)
			newPackCount[size]++

			heap.Push(pq, &state{
				totalItems: nextTotal,
				numPacks:   curr.numPacks + 1,
				packCount:  newPackCount,
			})
		}
	}

	// No solution found (should never happen with positive pack sizes)
	return nil
}

// TotalItems returns total number of items in packs
func TotalItems(packs []PackResult) int {
	total := 0
	for _, p := range packs {
		total += p.Size * p.Count
	}
	return total
}

// ---- internal types ----

// state represents a state in the priority queue
type state struct {
	totalItems int
	numPacks   int
	packCount  map[int]int
}

// priorityQueue implements heap.Interface based on totalItems, then numPacks
type priorityQueue []*state

// Len returns the number of items in the priority queue
func (pq priorityQueue) Len() int { return len(pq) }

// Less compares two states based on totalItems and numPacks
func (pq priorityQueue) Less(i, j int) bool {
	// Priority: smaller totalItems, then fewer packs
	if pq[i].totalItems != pq[j].totalItems {
		return pq[i].totalItems < pq[j].totalItems
	}
	return pq[i].numPacks < pq[j].numPacks
}

// Swap swaps two elements in the priority queue
func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

// Push adds a new state to the priority queue
func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*state))
}

// Pop removes and returns the last element from the priority queue
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// copyMap creates a shallow copy of a map[int]int
func copyMap(original map[int]int) map[int]int {
	newMap := make(map[int]int, len(original))
	for k, v := range original {
		newMap[k] = v
	}
	return newMap
}
