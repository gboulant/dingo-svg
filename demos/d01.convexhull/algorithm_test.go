package main

import (
	"testing"

	svg "github.com/gboulant/dingo-svg"
)

func TestAlgorithms_ConvexHull_Basic(t *testing.T) {
	var coords []struct{ X, Y float64 } = []struct{ X, Y float64 }{
		{0.2, 0.2},
		{0.3, 0.8},
		{0.6, 0.6},
		{0.8, 0.8},
		{0.8, 0.2},
	}
	points := *NewVectorsFromTuples(coords)
	hull, err := ConvexHull(points)
	if err != nil {
		t.Fatal(err)
	}

	expect := []int{0, 1, 3, 4}
	for i, e := range expect {
		if hull[i] != e {
			t.Errorf("the index %d is %d (should be %d)", i, hull[i], e)
		}
	}

}

func TestAlgorithms_ConvexHull_LargeDataset(t *testing.T) {
	points := dataset_random(0, 1, 0, 1, 16)
	hull, err := ConvexHull(points)
	if err != nil {
		t.Fatal(err)
	}

	// Draw the points and the convex hull using the SVG sketcher
	s := svg.NewSketcher()
	pencilblk := svg.NewPencil("black", 2)
	pencilred := svg.NewPencil("red", 2)

	s.Pencil = pencilblk
	for _, p := range points {
		s.Point(p.X, p.Y)
	}

	hullpolygon := make(Vectors, len(hull))
	for i, index := range hull {
		hullpolygon[i] = points[index]
	}
	s.Pencil = pencilred
	s.Polygon(hullpolygon.Tuples(), true)
	err = s.Save("output.TestAlgorithms_ConvexHull_LargeDataset.svg")
	if err != nil {
		t.Error(err)
	}
}
