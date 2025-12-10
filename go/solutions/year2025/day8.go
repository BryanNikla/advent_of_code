package year2025

import (
	"fmt"
	"sort"
	"strings"

	"advent_of_code/registry"
	"advent_of_code/utils"
)

func init() {
	registry.RegisterSolution(2025, 8, func() utils.Solution {
		input1, input2 := utils.GetInput(2025, 8)
		return utils.Solution{
			Day:   8,
			Test1: day8part1(input1) == 40,
			Test2: day8part2(input2) == 25272,
		}
	})
}

// circuitConnection represents a connection between two 3D coordinates and the distance between them
type circuitConnection struct {
	A        utils.Coordinate3D
	B        utils.Coordinate3D
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

func parseLinesIntoCoordinate3Ds(lines []string) []utils.Coordinate3D {
	all := make([]utils.Coordinate3D, len(lines))
	for i, line := range lines {
		coords := strings.Split(line, ",")
		all[i] = utils.Coordinate3D{
			ID: i,
			X:  utils.StringToFloat64(coords[0]),
			Y:  utils.StringToFloat64(coords[1]),
			Z:  utils.StringToFloat64(coords[2]),
		}
	}
	return all
}

// parseAllConnectionsMap generates all possible connections between points
func parseIntoConnectionsMap(all []utils.Coordinate3D) map[string]circuitConnection {
	allConnectionsMap := make(map[string]circuitConnection)
	// Generate all possible connections
	for _, a := range all {
		for _, b := range all {
			if a != b {
				// Create a unique key for the connection (order-independent); key is always lowerID-higherID
				id1, id2 := a.ID, b.ID
				if id1 > id2 {
					id1, id2 = id2, id1
				}
				key := fmt.Sprintf("%d-%d", id1, id2)

				// Add connection if it doesn't already exist
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

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func day8part1(input string) int {
	all := parseLinesIntoCoordinate3Ds(utils.GetLines(input))

	allConnectionsMap := parseIntoConnectionsMap(all) // Generate all possible connections

	allConnections := make([]circuitConnection, 0, len(allConnectionsMap))
	for _, conn := range allConnectionsMap {
		allConnections = append(allConnections, conn)
	}

	// Sort connections by distance (shortest to longest)
	sort.Slice(allConnections, func(i, j int) bool {
		return allConnections[i].Distance < allConnections[j].Distance
	})

	///////////////////////////////////////////////////////////////
	var connectionsToKeep = 10 // Change to 1000 for real input
	///////////////////////////////////////////////////////////////

	shortestConnections := allConnections[:connectionsToKeep]

	dsu := utils.NewDisjointSetUnion[int]()

	// Build
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

	var sizes []int
	for _, connections := range groups {
		sizes = append(sizes, countDistinct(connections))
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return utils.Multiply(sizes[0], sizes[1], sizes[2])
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func day8part2(input string) int {
	all := parseLinesIntoCoordinate3Ds(utils.GetLines(input))

	allConnectionsMap := parseIntoConnectionsMap(all) // Generate all possible connections

	allConnections := make([]circuitConnection, 0, len(allConnectionsMap))
	for _, conn := range allConnectionsMap {
		allConnections = append(allConnections, conn)
	}

	// Sort connections by distance (shortest to longest)
	sort.Slice(allConnections, func(i, j int) bool {
		return allConnections[i].Distance < allConnections[j].Distance
	})

	junctionBoxCount := len(all)

	// Add connections one by one until all junction boxes are connected
	// TODO: This is pretty inefficient (but works).. Could be optimized probably.
	for connCount := 0; connCount < len(allConnections); connCount++ {
		testSet := allConnections[:connCount]

		dsu := utils.NewDisjointSetUnion[int]()

		// Build relationships (Union)
		for _, item := range testSet {
			dsu.Union(item.A.ID, item.B.ID)
		}

		// Group the structs based on the Root of their IDs
		groups := make(map[int][]circuitConnection)

		for _, item := range testSet {
			// We can check A or B, they are in the same set now.
			root := dsu.Find(item.A.ID)
			groups[root] = append(groups[root], item)
		}

		groupsSlice := make([][]circuitConnection, 0, len(groups))
		for _, group := range groups {
			groupsSlice = append(groupsSlice, group)
		}

		if len(groupsSlice) != 1 {
			continue // Not all connected yet
		}

		// Check if this group contains all junction boxes now
		if countDistinct(groupsSlice[0]) == junctionBoxCount {
			lastConnection := testSet[len(testSet)-1]
			return int(utils.Multiply(lastConnection.A.X, lastConnection.B.X))
		}
	}

	fmt.Println("⚠️ Something is very wrong")
	return 0 // Should ideally never reach here.
}
