package year2025

import (
	"fmt"
	"sort"
	"strings"

	"advent_of_code/registry"
	"advent_of_code/utils"
)

func init() {
	registry.RegisterSolution(2025, 9, func() utils.Solution {
		input1, input2 := utils.GetInput(2025, 9)
		return utils.Solution{
			Day:   9,
			Test1: day9part1(input1) == 50,
			Test2: day9part2(input2) == 24,
		}
	})
}

func day9part1(input string) int {
	lines := utils.GetLines(input)

	coordinates := make([]utils.Coordinates, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		coordinates[i] = utils.Coordinates{
			ID: i,
			X:  utils.StringToInteger(parts[0]),
			Y:  utils.StringToInteger(parts[1]),
		}
	}

	allCoordinatePairs := parseIntoCoordinatePairs(coordinates)

	var largest int
	for _, pair := range allCoordinatePairs {
		if pair.Area > largest {
			largest = pair.Area
		}
	}
	return largest
}

func day9part2(input string) int {
	lines := utils.GetLines(input)

	coordinates := make([]utils.Coordinates, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		coordinates[i] = utils.Coordinates{
			ID: i,
			X:  utils.StringToInteger(parts[0]),
			Y:  utils.StringToInteger(parts[1]),
		}
	}

	allCoordinatePairs := parseIntoCoordinatePairs(coordinates)

	// Sort pairs by area in descending order
	sort.Slice(allCoordinatePairs, func(i, j int) bool {
		return allCoordinatePairs[i].Area > allCoordinatePairs[j].Area
	})

	for _, pair := range allCoordinatePairs {
		corners := FindCornersOfRectangleBetweenCoordinates(pair.A, pair.B)

		cornersAreGood := true
		for _, corner := range corners {
			if !IsPointInOrOnPolygon(corner, coordinates) {
				cornersAreGood = false
				break
			}
		}

		if cornersAreGood {
			if !RectangleIntersectsPolygonEdges(corners, coordinates) {
				fmt.Printf("ANSWER: %d\n", pair.Area)
				return pair.Area
			}

		}
	}

	fmt.Println("⚠️ Something went wrong, no valid rectangle found!")
	return 0
}

type CoordinatePairs struct {
	A    utils.Coordinates
	B    utils.Coordinates
	Area int
}

// parseAllConnectionsMap generates all possible connections between points
func parseIntoCoordinatePairs(all []utils.Coordinates) []CoordinatePairs {
	allCoordinatePairsMap := make(map[string]CoordinatePairs)
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
				if _, exists := allCoordinatePairsMap[key]; !exists {
					allCoordinatePairsMap[key] = CoordinatePairs{
						A:    a,
						B:    b,
						Area: FindAreaBetweenCoordinates(a, b),
					}
				}
			}
		}
	}
	// Convert map to slice
	coordinatePairsSlice := make([]CoordinatePairs, 0, len(allCoordinatePairsMap))
	for _, pair := range allCoordinatePairsMap {
		coordinatePairsSlice = append(coordinatePairsSlice, pair)
	}

	return coordinatePairsSlice
}

func FindAreaBetweenCoordinates(a, b utils.Coordinates) int {
	width := utils.AbsoluteValue(a.X-b.X) + 1
	height := utils.AbsoluteValue(a.Y-b.Y) + 1
	return width * height
}

func FindCornersOfRectangleBetweenCoordinates(a, b utils.Coordinates) []utils.Coordinates {
	minX := minInt(a.X, b.X)
	maxX := maxInt(a.X, b.X)
	minY := minInt(a.Y, b.Y)
	maxY := maxInt(a.Y, b.Y)

	return []utils.Coordinates{
		{X: minX, Y: minY}, // Bottom-left
		{X: maxX, Y: minY}, // Bottom-right
		{X: maxX, Y: maxY}, // Top-right
		{X: minX, Y: maxY}, // Top-left
	}
}

// IsPointInOrOnPolygon returns true if p is inside or on the boundary of the polygon
func IsPointInOrOnPolygon(p utils.Coordinates, polygon []utils.Coordinates) bool {
	inside := false
	n := len(polygon)
	j := n - 1 // The last vertex

	for i := 0; i < n; i++ {
		curr := polygon[i]
		prev := polygon[j]

		// If the point is exactly on the segment connecting curr and prev
		if isOnSegment(p, curr, prev) {
			return true
		}

		// Standard ray casting algorithm
		if (curr.Y > p.Y) != (prev.Y > p.Y) {
			intersectX := float64(prev.X-curr.X)*float64(p.Y-curr.Y)/
				float64(prev.Y-curr.Y) + float64(curr.X)

			if float64(p.X) < intersectX {
				inside = !inside
			}
		}

		j = i
	}

	return inside
}

// isOnSegment checks if point p lies on the line segment between a and b
func isOnSegment(p, a, b utils.Coordinates) bool {
	minX, maxX := minInt(a.X, b.X), maxInt(a.X, b.X)
	minY, maxY := minInt(a.Y, b.Y), maxInt(a.Y, b.Y)

	if p.X < minX || p.X > maxX || p.Y < minY || p.Y > maxY {
		return false
	}

	// Check collinearity using cross product
	// If (b.x - a.x) * (p.y - a.y) == (b.y - a.y) * (p.x - a.x), they are collinear.
	// This avoids division and floating point issues.
	crossProduct := (b.X-a.X)*(p.Y-a.Y) - (b.Y-a.Y)*(p.X-a.X)

	return crossProduct == 0
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func RectangleIntersectsPolygonEdges(rectCorners []utils.Coordinates, polygon []utils.Coordinates) bool {
	// Check every edge of the rectangle against every edge of the polygon
	n := len(polygon)
	for i := 0; i < 4; i++ {
		r1 := rectCorners[i]
		r2 := rectCorners[(i+1)%4] // Wrap around for rectangle edges

		for j := 0; j < n; j++ {
			p1 := polygon[j]
			p2 := polygon[(j+1)%n] // Wrap around for polygon edges

			if doSegmentsProperlyIntersect(r1, r2, p1, p2) {
				return true
			}
		}
	}
	return false
}

// Check if segment (p1,q1) properly intersects with segment (p2,q2).
// "Proper" means they cross each other strictly internally.
func doSegmentsProperlyIntersect(p1, q1, p2, q2 utils.Coordinates) bool {
	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)

	if o1 != 0 && o2 != 0 && o3 != 0 && o4 != 0 && o1 != o2 && o3 != o4 {
		return true
	}

	return false
}

// orientation finds the ordered triplet orientation of (p, q, r).
// Returns:
//   0 -> Collinear
//   1 -> Clockwise
//   2 -> Counterclockwise
func orientation(p, q, r utils.Coordinates) int {
	val := (q.Y-p.Y)*(r.X-q.X) - (q.X-p.X)*(r.Y-q.Y)

	if val == 0 {
		return 0
	}
	if val > 0 {
		return 1 // Clockwise
	}
	return 2 // Counterclockwise
}
