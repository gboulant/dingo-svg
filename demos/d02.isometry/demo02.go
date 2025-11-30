package main

func demo02() error {
	xyrange := 10. // the axis ranges (-xyrange/2, +xyrange/2)
	v := NewIsometricView(xyrange)
	v.DrawAxis()

	lx := 4.
	ly := 3.
	lz := 5.
	fill := false

	// zface is the horizontal face at z = lz/2
	var zface [][3]float64 = [][3]float64{
		{-lx / 2, -ly / 2, lz / 2},
		{-lx / 2, +ly / 2, lz / 2},
		{+lx / 2, +ly / 2, lz / 2},
		{+lx / 2, -ly / 2, lz / 2},
	}
	v.DrawPolygon(zface, fill)

	// xface is the vertical face at x=lx/2
	var xface [][3]float64 = [][3]float64{
		{+lx / 2, -ly / 2, -lz / 2},
		{+lx / 2, +ly / 2, -lz / 2},
		{+lx / 2, +ly / 2, +lz / 2},
		{+lx / 2, -ly / 2, +lz / 2},
	}
	v.DrawPolygon(xface, fill)

	// yface is the vertical face at y=ly/2
	var yface [][3]float64 = [][3]float64{
		{-lx / 2, +ly / 2, -lz / 2},
		{+lx / 2, +ly / 2, -lz / 2},
		{+lx / 2, +ly / 2, +lz / 2},
		{-lx / 2, +ly / 2, +lz / 2},
	}
	v.DrawPolygon(yface, fill)

	return v.Save("output.demo02.svg")
}
