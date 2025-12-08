package utils

// Disjoint Set Union (DSU) or Union-Find data structure implementation
// Used to efficiently manage and merge disjoint sets
// Supports two primary operations: Find and Union
// Find: Determine the representative (root) of the set containing a specific element
// Union: Merge two sets into a single set
//
// Interesting read: https://www.geeksforgeeks.org/dsa/introduction-to-disjoint-set-data-structure-or-union-find-algorithm/
// Useful for problems involving connectivity, clustering, and network components

type DSU struct {
	// Parent map to track the root of each element
	parent map[int]int // key: element ID, value: parent ID
}

func NewDisjointSetUnion() *DSU {
	return &DSU{
		parent: make(map[int]int),
	}
}

// Find returns the representative (root) of the set 'i' belongs to.
func (d *DSU) Find(i int) int {
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
func (d *DSU) Union(i int, j int) {
	rootI := d.Find(i)
	rootJ := d.Find(j)
	if rootI != rootJ {
		d.parent[rootI] = rootJ
	}
}
