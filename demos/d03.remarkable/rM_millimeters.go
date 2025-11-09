package main

import svg "github.com/gboulant/dingo-svg"

func demo_rM_millimeters() error {
	s := NewRemarkableSketcher()

	// Prepare the pencils
	pthin := svg.NewPencil("Gray", 1)
	pbold := svg.NewPencil("Gray", 2)

	xmin, xmax, ymin, ymax := s.CoordinatesSystem().UserCoordinatesBoundaries()

	cellsize := 10. // mm
	nbtics := 5
	ticsize := float64(cellsize) / float64(nbtics)

	offset := 10. // mm (minimal margin offset)
	xrange := xmax - xmin
	xnbcells := int(float64(xrange-offset) / cellsize)
	xmargin := 0.5 * (xrange - float64(xnbcells)*cellsize)

	yrange := ymax - ymin
	ynbcells := int(float64(yrange-offset) / cellsize)
	ymargin := 0.5 * (yrange - float64(ynbcells)*cellsize)

	// ------------------------------------------------
	// millimeters grid
	s.Pencil = pthin
	x := xmargin
	for x <= xmax-xmargin {
		s.MoveTo(x, ymin+ymargin)
		s.LineTo(x, ymax-ymargin)
		x += ticsize

	}
	y := ymargin
	for y <= ymax-ymargin {
		s.MoveTo(xmin+xmargin, y)
		s.LineTo(xmax-xmargin, y)
		y += ticsize
	}

	// ------------------------------------------------
	// centimeters grid
	s.Pencil = pbold
	x = xmargin
	for x <= xmax-xmargin {
		s.MoveTo(x, ymin+ymargin)
		s.LineTo(x, ymax-ymargin)
		x += cellsize
	}
	y = ymargin
	for y <= ymax-ymargin {
		s.MoveTo(xmin+xmargin, y)
		s.LineTo(xmax-xmargin, y)
		y += cellsize
	}

	// scale legend
	s.Pencil.FontSize = 20
	s.Edge(xmargin, ymargin-1, xmargin+cellsize, ymargin-1)
	s.Text(xmargin+3, ymargin-3, "1 cm")

	s.Save("output.demo_rM_millimeters.svg")
	return nil
}
