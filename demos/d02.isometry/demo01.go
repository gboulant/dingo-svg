package main

/*
This example is inspired from the example ch3/surface of the book "The Go
Programing Language" from Donovan and Ritchie. See source code at:

	https://github.com/adonovan/gopl.io

*/

func demo01_cardinalsine() error {

	xymax := 30.
	period := xymax / 4.     // spatial period of the cardinal sine
	amplitude := 0.4 * xymax // amplitude of the cardianl sine
	f := CardinalSine(period, amplitude)

	gridsize := 80
	v := DrawIsometricView(f, gridsize, xymax)
	return v.Save("output.demo01.cardinalsine.svg")
}

func demo01_horseshoe() error {

	xymax := 2.
	f := func(x, y float64) float64 {
		return (x*x - y*y)
	}

	gridsize := 80
	v := DrawIsometricView(f, gridsize, xymax)
	return v.Save("output.demo01.horseshoe.svg")
}

func demo01_parabol() error {

	xymax := 2.
	f := func(x, y float64) float64 {
		return (x*x + y*y)
	}

	gridsize := 80
	v := DrawIsometricView(f, gridsize, xymax)
	return v.Save("output.demo01.parabol.svg")
}
