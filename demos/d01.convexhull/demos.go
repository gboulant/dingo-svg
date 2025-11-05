package main

import (
	"io"
	"log"

	svg "github.com/gboulant/dingo-svg"
)

func init() {
	log.SetOutput(io.Discard)
}

func demo01() error {
	points := dataset_random(0, 1, 0, 1, 16)
	hull, err := ConvexHull(points)
	if err != nil {
		return err
	}

	// Draw the points and the convex hull using the SVG sketcher. We use
	// different pencils for drawing with different colors
	s := svg.NewSketcher()
	pencilblk := svg.NewPencil("black", 2)
	pencilred := svg.NewPencil("red", 2)

	// Draw the points in black
	s.Pencil = pencilblk
	for _, p := range points {
		s.Point(p.X, p.Y)
	}

	// Draw the hull polygon in red
	hullpolygon := make(Vectors, len(hull))
	for i, index := range hull {
		hullpolygon[i] = points[index]
	}
	s.Pencil = pencilred
	s.Polygon(hullpolygon.Tuples(), true)

	err = s.Save("output.demo01.svg")
	if err != nil {
		return err
	}
	return nil
}
