package main

import (
	"fmt"

	svg "github.com/gboulant/dingo-svg"
)

func hline(sk *svg.Sketcher, y float64, xmin, xmax float64) {
	sk.MoveTo(xmin, y)
	sk.LineTo(xmax, y)
}

func vline(sk *svg.Sketcher, x float64, ymin, ymax float64) {
	sk.MoveTo(x, ymin)
	sk.LineTo(x, ymax)
}

func makesketch(notes Notes, svgpath string) error {
	nbstrings := len(notes)
	nbfrets := len(notes[0])

	cnvwidth := 1280
	cnvheight := cnvwidth * nbstrings / nbfrets

	// The choice of the cell size is in fact arbitrary. It only defined
	//the xrange to consider and it has no effect on the size of the
	//element because all positions are computed with this unit size.
	xcellsize := 1.
	xrange := xcellsize * float64(nbfrets+1)
	recsize := xcellsize * 0.6

	cs := svg.NewCoordSysTopLeft(cnvwidth, cnvheight, xrange)
	sk := svg.NewSketcher().WithCoordinateSystem(cs).WithBackgroundColor("white")
	sk.Pencil.FontFamily = "monospace"
	sk.Pencil.FontSize = 18

	xmin, xmax, ymin, ymax := sk.CoordinatesSystem().UserCoordinatesBoundaries()

	var ys float64
	var xf float64

	sk.Pencil.LineColor = "lightgray"
	sk.Pencil.FillColor = "white"
	stringNotes := notes[0]
	ys = xcellsize * 0.3
	sk.Pencil.FontWeight = "bold"
	sk.Pencil.FontColor = "orange"
	for j := range nbfrets {
		note := stringNotes[j]
		xf = xcellsize * float64(note.FretNumber+1)

		sk.Pencil.LineWidth = 1
		vline(sk, xf, ymin, ymax)

		sk.Pencil.LineWidth = 0
		sk.Rectangle(xf-recsize*0.5, ys-recsize*0.5, recsize, recsize*0.8, true)
		sk.Text(xf-recsize*0.4, ys+recsize*0.1, fmt.Sprintf("F%.2d", note.FretNumber))

	}

	sk.Pencil.FontColor = svg.DefaultFontColor
	sk.Pencil.FontWeight = svg.DefaultFontWeight
	for i := range nbstrings {
		stringNotes := notes[i]
		stringNumber := stringNotes[0].StringNumber
		ys = xcellsize * float64(stringNumber)

		sk.Pencil.LineWidth = 1
		hline(sk, ys, xmin, xmax)

		sk.Pencil.LineWidth = 0
		sk.Pencil.FontWeight = "bold"
		sk.Pencil.FontColor = "orange"
		xf = xcellsize * 0.3
		sk.Rectangle(xf-recsize*0.5, ys-recsize*0.5, recsize, recsize*0.8, true)

		sk.Pencil.FontSize = 18
		sk.Text(xf-recsize*0.4, ys+recsize*0.1, fmt.Sprintf("S%d", stringNumber))

		sk.Pencil.FontColor = svg.DefaultFontColor
		sk.Pencil.FontWeight = svg.DefaultFontWeight
		for j := range nbfrets {
			note := stringNotes[j]
			xf = xcellsize * float64(note.FretNumber+1)
			sk.Rectangle(xf-recsize*0.5, ys-recsize*0.5, recsize, recsize*0.8, true)

			sk.Pencil.FontSize = 18
			sk.Text(xf-recsize*0.4, ys+recsize*0.1, note.Name)
			sk.Pencil.FontSize = 12
			sk.Text(xf-recsize*0.4, ys+recsize*0.4, note.Frequency)
		}
	}

	return sk.Save(svgpath)
}
