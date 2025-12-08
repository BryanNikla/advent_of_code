package utils

import (
	"math"
)

type Coordinate3D struct {
	ID int // Unique identifier for the coordinate point (helps in distinguishing points, ignore otherwise)
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
