package main

import (
	"fmt"
	"math"
	"testing"
)

func TestVector_AngleTo01(t *testing.T) {
	u := NewVector(1.0, 0.0)
	v := NewVector(0.5, 0.5)

	arad := u.AngleTo(*v)
	adeg := 180 * arad / math.Pi
	fmt.Printf("angle(u->v) = %.2f\n", adeg)
	aref := 45.0
	if !AlmostEqual(adeg, aref, epsilon) {
		t.Errorf("angle is %.2f (should be %.2f)", adeg, aref)
	}

	arad = v.AngleTo(*u)
	adeg = 180 * arad / math.Pi
	fmt.Printf("angle(v->u) = %.2f\n", adeg)
	aref = 360. - 45.
	if !AlmostEqual(adeg, aref, epsilon) {
		t.Errorf("angle is %.2f (should be %.2f)", adeg, aref)
	}
}

func TestVector_AngleTo02(t *testing.T) {

	var coords []struct{ X, Y float64 } = []struct{ X, Y float64 }{
		{0.2, 0.2},
		{0.3, 0.8},
		{0.6, 0.6},
		{0.8, 0.8},
		{0.8, 0.2},
	}
	points := *NewVectorsFromTuples(coords)
	refpoint := points[0]
	others := points[1:]

	var angles []float64 = make([]float64, len(others))
	v := Vector{X: 0.0, Y: 1.0}
	for i, p := range others {
		u := refpoint.VectorTo(p)
		a := u.AngleTo(v)
		fmt.Printf("p%d (%s): a=%.2f\n", i, p, Rad2Deg(a))
		angles[i] = Rad2Deg(a)
	}

	// The first angle should be around 10°, then around 45, and the last 90
	expect := []float64{9.46232221, 45.0, 45.0, 90.0}
	for i, a := range expect {
		if !AlmostEqual(angles[i], a, 1e-8) {
			t.Errorf("angle is %.8f (should be %.8f)", angles[i], a)
		}
	}
}

func TestVector_RemoveElement(t *testing.T) {

	var coords []struct{ X, Y float64 } = []struct{ X, Y float64 }{
		{0.2, 0.2},
		{0.3, 0.8},
		{0.6, 0.6},
		{0.8, 0.8},
		{0.8, 0.2},
	}
	points := *NewVectorsFromTuples(coords)
	len0 := len(points)

	v := points.RemoveByShift(2)
	if v.X != 0.6 || v.Y != 0.6 {
		t.Errorf("the remove point is %s (should be [0.6, 0.6])", v)
	}
	if len(points) != len0-1 {
		t.Errorf("the new points list has len %d (should be %d)", len(points), len0)
	}
	expect := []struct{ X, Y float64 }{
		{0.2, 0.2},
		{0.3, 0.8},
		{0.8, 0.8},
		{0.8, 0.2},
	}

	for i, e := range expect {
		if points[i].Tuple() != e {
			t.Errorf("the coords are %s (should be %v)", points[i], e)
		}
	}
}
