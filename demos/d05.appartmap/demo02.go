package main

import (
	"math"

	svg "github.com/gboulant/dingo-svg"
)

type Wall struct {
	X1, Y1 float64
	X2, Y2 float64
}

func (w Wall) Length() float64 {
	ux := w.X2 - w.X1
	uy := w.Y2 - w.Y1
	return math.Sqrt(ux*ux + uy*uy)
}

func (w Wall) NextLeft(l float64) Wall {
	wl := w.Length()
	ux := (w.X2 - w.X1) / wl
	uy := (w.Y2 - w.Y1) / wl
	vx := -uy
	vy := ux

	X1 := w.X2
	Y1 := w.Y2
	X2 := X1 + l*vx
	Y2 := Y1 + l*vy

	return Wall{X1, Y1, X2, Y2}
}

type Space struct {
	Walls []Wall
}

func NewSpace(x, y float64, w, h float64) *Space {
	return &Space{Walls: []Wall{
		{X1: x, Y1: y, X2: x + w, Y2: y},
		{X1: x + w, Y1: y, X2: x + w, Y2: y + h},
		{X1: x + w, Y1: y + h, X2: x, Y2: y + h},
		{X1: x, Y1: y + h, X2: x, Y2: y},
	}}
}

func (s Space) Draw(sk *svg.Sketcher) {
	for _, w := range s.Walls {
		sk.Edge(w.X1, w.Y1, w.X2, w.Y2)
	}
}

func demo02() {

	cs := svg.NewCoordSysBottomLeft(cnvwidth, cnvheight, xrange)
	sk := svg.NewSketcher().WithCoordinateSystem(cs)

	W := 40.
	H := 25.
	X := 10.
	Y := 10.

	x := X
	y := Y
	w := W
	h := H
	sp := NewSpace(x, y, w, h)
	sp.Draw(sk)

	x = X + W
	y = Y
	w = H
	h = W
	walls := make([]Wall, 4)
	walls[0] = Wall{x, y, x + w, y}
	walls[1] = walls[0].NextLeft(h)
	walls[2] = walls[1].NextLeft(w)
	walls[3] = walls[2].NextLeft(h)
	sp = &Space{Walls: walls}
	sp.Draw(sk)

	sk.Save("output.demo02.svg")
}
