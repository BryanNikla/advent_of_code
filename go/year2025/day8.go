package year2025

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"advent_of_code/registry"
	"advent_of_code/utils"
)

func init() {
	registry.RegisterSolution(2025, 8, func() utils.Solution {
		input1, input2 := utils.GetInput(2025, 8)
		return utils.Solution{
			Day:   7,
			Test1: day8part1(input1) == 40,
			Test2: day8part2(input2) == 25272,
		}
	})
}

type Coordinate3D struct {
	ID int
	X  float64
	Y  float64
	Z  float64
}

// DistanceTo calculates the Euclidean distance between point p and point q
// sqrt(dx^2 + dy^2 + dz^2)
func (p Coordinate3D) DistanceTo(q Coordinate3D) float64 {
	dx := p.X - q.X
	dy := p.Y - q.Y
	dz := p.Z - q.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

type circuitConnection struct {
	ID       int
	A        Coordinate3D
	B        Coordinate3D
	Distance float64
}

// countDistinct calculates the number of unique points in a connection list
func countDistinct(group []circuitConnection) int {
	distinct := make(map[int]bool)
	for _, conn := range group {
		distinct[conn.A.ID] = true
		distinct[conn.B.ID] = true
	}
	return len(distinct)
}

// DSU ( Disjoint Set Union )
type DSU struct {
	parent map[int]int
}

func NewDSU() *DSU {
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
	d.parent[i] = d.Find(d.parent[i]) // Path compression: point directly to root
	return d.parent[i]
}

// Union merges the sets containing i and j.
func (d *DSU) Union(i int, j int) {
	rootI := d.Find(i)
	rootJ := d.Find(j)
	if rootI != rootJ {
		d.parent[rootI] = rootJ
	}
}

func parseLinesIntoCoordinate3Ds(lines []string) []Coordinate3D {
	all := make([]Coordinate3D, len(lines))
	for i, line := range lines {
		coords := strings.Split(line, ",")
		all[i] = Coordinate3D{
			ID: i,
			X:  utils.StringToFloat64(coords[0]),
			Y:  utils.StringToFloat64(coords[1]),
			Z:  utils.StringToFloat64(coords[2]),
		}
	}
	return all
}

// parseAllConnectionsMap generates all possible connections between points
func parseIntoConnectionsMap(all []Coordinate3D) map[string]circuitConnection {
	allConnectionsMap := make(map[string]circuitConnection)
	// Generate all possible connections
	for _, a := range all {
		for _, b := range all {
			if a != b {
				id1, id2 := a.ID, b.ID
				if id1 > id2 {
					id1, id2 = id2, id1
				}
				key := fmt.Sprintf("%d-%d", id1, id2)
				if _, exists := allConnectionsMap[key]; !exists {
					allConnectionsMap[key] = circuitConnection{
						A:        a,
						B:        b,
						Distance: a.DistanceTo(b),
					}
				}
			}
		}
	}
	return allConnectionsMap
}

func day8part1(input string) int {
	all := parseLinesIntoCoordinate3Ds(utils.GetLines(input))

	// Generate all possible connections
	allConnectionsMap := parseIntoConnectionsMap(all)

	allConnections := make([]circuitConnection, 0, len(allConnectionsMap))
	for _, conn := range allConnectionsMap {
		allConnections = append(allConnections, conn)
	}

	// Sort connections by distance (shortest to longest)
	sort.Slice(allConnections, func(i, j int) bool {
		return allConnections[i].Distance < allConnections[j].Distance
	})

	///////////////////////////////////////////////////////////////
	// IMPORTANT VARIABLE TO CHANGE FOR REAL RUN
	var connectionsToKeep = 10 // test is 10
	//var connectionsToKeep = 1000 // Real input is 1000
	///////////////////////////////////////////////////////////////

	shortestConnections := allConnections[:connectionsToKeep]

	dsu := NewDSU()

	// Build relationships (Union)
	for _, item := range shortestConnections {
		dsu.Union(item.A.ID, item.B.ID)
	}

	// Group the structs based on the Root of their IDs
	groups := make(map[int][]circuitConnection)

	for _, item := range shortestConnections {
		// We can check A or B, they are in the same set now.
		root := dsu.Find(item.A.ID)
		groups[root] = append(groups[root], item)
	}

	groupsSlice := make([][]circuitConnection, 0, len(groups))
	for _, group := range groups {
		groupsSlice = append(groupsSlice, group)
	}

	// Struct to hold the group and its size
	type GroupInfo struct {
		Connections []circuitConnection
		Size        int
	}

	// Iterate over the map, determine circuit size, then add to slice
	var sortedGroups []GroupInfo
	for _, connections := range groups {
		sortedGroups = append(sortedGroups, GroupInfo{
			Connections: connections,
			Size:        countDistinct(connections),
		})
	}

	// Sort by circuit sizes (Largest to Smallest)
	sort.Slice(sortedGroups, func(i, j int) bool {
		return sortedGroups[i].Size > sortedGroups[j].Size
	})

	// Only keep the 3 largest
	threeLargest := sortedGroups[:3]

	result := 0
	for _, grp := range threeLargest {
		if result == 0 {
			result = grp.Size // Initialize math
		} else {
			result *= grp.Size
		}
	}

	return result
}

func day8part2(input string) int {
	all := parseLinesIntoCoordinate3Ds(utils.GetLines(input))

	return 0
}
