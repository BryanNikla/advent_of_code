package utils

// DSU ( Disjoint Set Union )
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
