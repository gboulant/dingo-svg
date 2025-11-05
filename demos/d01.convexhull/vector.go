package main

import (
	"fmt"
	"math"
)

const TwoPi float64 = 2 * math.Pi

// ----------------------------------------------------------
// Vector definition
// ----------------------------------------------------------

// Vector defines a point by its coordinates in a 2D plane
type Vector struct {
	X float64
	Y float64
}

func NewVector(x, y float64) *Vector {
	return &Vector{X: x, Y: y}
}

func (u Vector) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", u.X, u.Y)
}

func (u Vector) Tuple() struct{ X, Y float64 } {
	return struct{ X, Y float64 }{
		X: u.X, Y: u.Y,
	}
}

func (u Vector) Norm() float64 {
	return math.Sqrt(u.X*u.X + u.Y*u.Y)
}

func (u Vector) ScalarProduct(v Vector) float64 {
	return ScalarProduct(u, v)
}

func (u Vector) VectorProductZ(v Vector) (z float64) {
	return VectorProductZ(u, v)
}

// VectorTo returns the vector that joins the point u to the point v,
// i.e. the vector w such that u + w = v, then w = v - u
func (u Vector) VectorTo(v Vector) Vector {
	return VectorSub(v, u)
}

func (u Vector) AngleTo(v Vector) float64 {
	norm := u.Norm() * v.Norm()
	cosa := u.ScalarProduct(v) / norm
	sina := u.VectorProductZ(v) / norm
	angle := math.Acos(cosa)
	if sina < 0 {
		angle = TwoPi - angle
	}
	return angle
}

// ----------------------------------------------------------
// Vectors opération
// ----------------------------------------------------------
// VectorSub returns  the vector u-v
func VectorSub(u, v Vector) Vector {
	return Vector{u.X - v.X, u.Y - v.Y}
}

// VectorAdd returns  the vector u+v
func VectorAdd(u, v Vector) Vector {
	return Vector{u.X + v.X, u.Y + v.Y}
}

func ScalarProduct(u, v Vector) float64 {
	return u.X*v.X + u.Y*v.Y
}

func VectorProductZ(u, v Vector) (z float64) {
	return u.X*v.Y - u.Y*v.X
}

// ----------------------------------------------------------
// Vectors definition
// ----------------------------------------------------------
type Vectors []Vector

func (vs Vectors) Tuples() []struct{ X, Y float64 } {
	tuples := make([]struct{ X, Y float64 }, len(vs))
	for i, v := range vs {
		tuples[i] = struct{ X, Y float64 }{v.X, v.Y}
	}
	return tuples
}

func NewVectorsFromTuples(tuples []struct{ X, Y float64 }) *Vectors {
	vectors := make(Vectors, len(tuples))
	for i, t := range tuples {
		vectors[i] = Vector{t.X, t.Y}
	}
	return &vectors
}

func (vs *Vectors) RemoveByShift(index int) Vector {
	v := (*vs)[index] // the element to be removed
	*vs = append((*vs)[:index], (*vs)[index+1:]...)
	return v
}

func (vs *Vectors) RemoveBySwitch(index int) Vector {
	v := (*vs)[index] // the element to be removed
	*vs = append((*vs)[:index], (*vs)[index+1:]...)
	return v
}
