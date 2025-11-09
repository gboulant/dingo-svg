package svg

import (
	"log"
	"testing"
)

func testpoints() []struct{ X, Y float64 } {
	var points []struct{ X, Y float64 } = []struct{ X, Y float64 }{
		{0.2, 0.2},
		{0.3, 0.8},
		{0.6, 0.6},
		{0.8, 0.8},
		{0.8, 0.2},
	}
	return points
}

func TestCoordSys_ChangeCoordinates(t *testing.T) {
	cnvwidth := DefaultCanvasWidth
	cnvheight := DefaultCanvasHeight
	xrange := 1.
	cs := NewCoordSysCentered(cnvwidth, cnvheight, xrange)

	Opx, Opy := cs.canvasCoordinates(0, 0)
	if Opx != 0.5*float64(cnvwidth) {
		t.Errorf("Opx is %.2f (should be %.2f)", Opx, 0.5*float64(cnvwidth))
	}
	if Opy != 0.5*float64(cnvheight) {
		t.Errorf("Opy is %.2f (should be %.2f)", Opy, 0.5*float64(cnvheight))
	}

	Mpx, Mpy := 0., float64(cnvheight)
	Mx, My := cs.userCoordinates(Mpx, Mpy)
	if Mx != -0.5 {
		t.Errorf("Mx is %.2f (should be %.2f)", Mx, -0.5)
	}
	if My != -0.5 {
		t.Errorf("My is %.2f (should be %.2f)", My, -0.5)
	}
}

func TestCoordSys_ImplementationOfWithRanges(t *testing.T) {
	points := testpoints()

	// ------------------------------------------------
	// Part 01: this first part shows the default implementation for
	// setting of the WithRange Coordinates System. It consists in determining
	// the canvas coordinates of the origin on the basis of the point xmin, ymin
	// that should correspond to the canvas coordinates (0, cnvheight), i.e. the
	// bottom left corner
	cnvwidth := DefaultCanvasWidth
	xmin, ymin, xmax, ymax := boundingBox(points)
	xrange := (xmax - xmin)
	yrange := (ymax - ymin)
	cnvheight := int((yrange / xrange) * float64(cnvwidth))

	unit2pixel := float64(cnvwidth) / xrange

	xsign := +1.
	ysign := -1.
	pxOrigin := 0.
	pyOrigin := float64(cnvheight)
	cnvXorigin := pxOrigin - xsign*unit2pixel*xmin
	cnvYorigin := pyOrigin - ysign*unit2pixel*ymin

	cs1 := newCoordSysAtOrigin(cnvXorigin, cnvYorigin, cnvwidth, cnvheight, xrange)
	pxmin, pymin := cs1.canvasCoordinates(xmin, ymin)
	if pxmin != pxOrigin {
		t.Errorf("pxmin is %.2f (should be %.2f)", pxmin, pxOrigin)
	}
	if pymin != pyOrigin {
		t.Errorf("pymin is %.2f (should be %.2f)", pymin, pyOrigin)
	}
	log.Println(cs1)

	// The problem with this implementation is that we have to redefine the
	// scaling factors unit2pixel and the axis orientation xsign and ysign prior
	// to the call of the function newCoordSysAtOrigin, while they are set in
	// the function itself. At best, we code twice the things, at worst we code
	// the things with different ways, which could inplies some conflicts

	// ------------------------------------------------
	// Part 02: this second part show an alternative implementation. It consists
	// in creating first a Coorinates System with origin at (0,0), so that we
	// can use the internal scaling functions of the coordinate system to reset
	// correctly the origin.
	cs2 := newCoordSysAtOrigin(0, 0, cnvwidth, cnvheight, xrange)
	// Compute the pixel position of the point xmin,ymin in this coordinate system
	pxmin, pymin = cs2.canvasCoordinates(xmin, ymin)
	// The real origin pixel position can be recomputed with
	cnvXorigin = pxOrigin - pxmin
	cnvYorigin = pyOrigin - pymin
	// We can then reset the origin of the coordinate system
	cs2.xorigin = cnvXorigin
	cs2.yorigin = cnvYorigin

	log.Println(cs2)

}

func TestCoordSysWithRanges(t *testing.T) {
	points := testpoints()
	cnvwidth := DefaultCanvasWidth
	xmin, ymin, xmax, ymax := boundingBox(points)
	cs := NewCoordSysWithRanges(cnvwidth, xmin, ymin, xmax, ymax)

	cnvheight := int(((ymax - ymin) / (xmax - xmin)) * float64(cnvwidth))

	pxmin, pymin := cs.canvasCoordinates(xmin, ymin)
	if pxmin != 0. {
		t.Errorf("pxmin is %.2f (should be %.2f)", pxmin, 0.)
	}
	if pymin != float64(cnvheight) {
		t.Errorf("pymin is %.2f (should be %.2f)", pymin, float64(cnvheight))
	}

	pxmax, pymax := cs.canvasCoordinates(xmax, ymax)
	if pxmax != float64(cnvwidth) {
		t.Errorf("pxmax is %.2f (should be %.2f)", pxmax, float64(cnvwidth))
	}
	if pymax != 0. {
		t.Errorf("pymax is %.2f (should be %.2f)", pymax, 0.)
	}

}
