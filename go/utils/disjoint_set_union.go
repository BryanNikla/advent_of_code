package utils

// Disjoint Set Union (DSU) or Union-Find data structure implementation
// Used to efficiently manage and merge disjoint sets
// Supports two primary operations: Find and Union
// Find: Determine the representative (root) of the set containing a specific element
// Union: Merge two sets into a single set
//
// Interesting read: https://www.geeksforgeeks.org/dsa/introduction-to-disjoint-set-data-structure-or-union-find-algorithm/
// Useful for problems involving connectivity, clustering, and network components

type DSU[T comparable] struct {
	// Parent map to track the root of each element
	parent map[T]T // key: element ID, value: parent ID
}

func NewDisjointSetUnion[T comparable]() *DSU[T] {
	return &DSU[T]{
		// 3. Use T for the map types instead of hardcoding int
		parent: make(map[T]T),
	}
}

// Find returns the representative (root) of the set 'i' belongs to.
func (d *DSU[T]) Find(i T) T {
	if _, exists := d.parent[i]; !exists {
		d.parent[i] = i
		return i
	}
	if d.parent[i] == i {
		return i
	}
	d.parent[i] = d.Find(d.parent[i])
	return d.parent[i]
}

// Union merges the sets containing i and j.
// If they are already in the same set, do nothing.
func (d *DSU[T]) Union(i T, j T) {
	rootI := d.Find(i)
	rootJ := d.Find(j)
	if rootI != rootJ {
		d.parent[rootI] = rootJ
	}
}
