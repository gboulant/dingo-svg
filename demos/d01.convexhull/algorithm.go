package main

import (
	"fmt"
	"log"
	"math"
)

// ----------------------------------------------------------
// Conversions

const epsilon float64 = 1e-12

func AlmostEqual(a, b float64, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

func Rad2Deg(angle_radian float64) float64 {
	return 180. * angle_radian / math.Pi
}

func Deg2Rad(angle_degree float64) float64 {
	return math.Pi * angle_degree / 180.
}

// ----------------------------------------------------------
// Convex hull returns the ordered list of indeces of the points that defined
// the convex hull of the input points. The points corresponding to these
// indeces define a polygon that corresponds to the convex hull of the input
// points.
func ConvexHull(points []Vector) ([]int, error) {
	hull := make([]int, 0)

	// -------------------------------------------
	// 1. Find the point with the minimal X value
	startindex := 0
	startpoint := points[startindex]
	for i, p := range points {
		if p.X < startpoint.X {
			startpoint = p
			startindex = i
		}
	}

	// Add this first point index in the hull. It is the starting point
	hull = append(hull, startindex)

	// -------------------------------------------
	// 2. Find next points at minimal angle from a reference vector. We start
	// with a vertical (Oy) reference vector becaus the starting point is the
	// point with minimal X value (and then can not have points on the left of
	// the vertical axis)
	refvector := Vector{X: 0.0, Y: 1.0}
	refpoint := startpoint
	refindex := startindex

	// The loop continue until we reach the starting point (refindex ==
	// startindex), or if the number of iterations exceeds the number of points
	// in this case, we delare an error)
	var err error = nil
	maxloop := len(points)
	nbloops := 0

	ended := false
	for !ended {
		amin := TwoPi
		idxpoint := 0
		for i, p := range points {
			if i == refindex {
				continue
			}
			u := refpoint.VectorTo(p) // vector joining refpoint to p
			a := u.AngleTo(refvector) // angle to the reference vector
			log.Printf("p%d at %s: a=%.2f\n", i, p, Rad2Deg(a))
			if a < amin {
				amin = a
				idxpoint = i
			}
		}
		log.Printf("minimal angle: p%d at %s: amin=%.2f\n", idxpoint, points[idxpoint], Rad2Deg(amin))
		refindex = idxpoint

		if refindex == startindex {
			ended = true
		} else {
			lastpoint := refpoint
			refpoint = points[refindex]
			refvector = lastpoint.VectorTo(refpoint)
			hull = append(hull, refindex)
		}

		nbloops += 1
		if !ended && nbloops > maxloop {
			ended = true
			err = fmt.Errorf("the hull search loop number (%d) exceeds the number of points (%d)", nbloops, maxloop)
		}
	}
	if err != nil {
		return nil, fmt.Errorf("the convex hull search fails du to error: %s", err)
	}
	return hull, nil
}
