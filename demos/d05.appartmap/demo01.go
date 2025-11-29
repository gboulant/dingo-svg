package main

import svg "github.com/gboulant/dingo-svg"

type Rectangle struct {
	X0, Y0        float64
	Width, Height float64
}

func (r Rectangle) Draw(sk *svg.Sketcher) {
	sk.Rectangle(r.X0, r.Y0, r.Width, r.Height, false)
}

func demo01() {

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
	r := Rectangle{x, y, w, h}
	r.Draw(sk)

	x = x + w
	y = y + 0.2*h
	w = 1.2 * H
	h = 0.8 * W
	r = Rectangle{x, y, w, h}
	r.Draw(sk)

	w = 1.2 * W
	h = 1.0 * H
	x = x - w
	y = Y + H
	r = Rectangle{x, y, w, h}
	r.Draw(sk)

	sk.Save("output.demo01.svg")
}
