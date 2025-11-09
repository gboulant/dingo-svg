package svg

import (
	"fmt"
	"math"
)

// ===========================================================================
// Coordinates System
// ===========================================================================

const (
	DefaultCanvasWidth  = 600
	DefaultCanvasHeight = 600
)

type CoordinateSystem struct {
	cnvxsize   int     // canvas width in pixels
	cnvysize   int     // canvas height in pixels
	xorigin    float64 // position of the origin on the X axis in pixels
	yorigin    float64 // position of the origin on the Y axis in pixels
	xsign      float64 // orientation of the X axis: -1 means from right to left
	ysign      float64 // orientation of the Y axis: -1 means from bottom to top
	unit2pixel float64 // number of pixels in a user unit (could be a float)
}

func (c CoordinateSystem) String() string {
	return fmt.Sprintf("cnvsize: w=%d x h=%d, origin: Ox=%.2fpx Oy=%.2fpx",
		c.cnvxsize, c.cnvysize, c.xorigin, c.yorigin)
}

// NewCoordinateSystem returns the default coordinates system. It is a
// coordinates system 1/ based on a canvas Width x Height, 2/ with origin at the
// bottom left corner, 3/ with user y axis oriented bottom up (inverse of the
// canvas native axis), and 4/ with an xrange equal to 1, i.e. x=1 corresponds
// to a x coordinates at the right boundary of the canvas (x=1 => xpixels =
// cnvwidth)
func NewCoordinateSystem() *CoordinateSystem {
	cnvxsize := DefaultCanvasWidth
	cnvysize := DefaultCanvasHeight
	xrange := 1. // coordinates are supposed to be in [0..1]
	return NewCoordSysBottomLeft(cnvxsize, cnvysize, xrange)
}

// canvasCoordinates returns the position of the point in the canvas native
// coordinates system, i.e. number of pixels from the top left corner (the
// native origin of the canvas) along the horizontal axis (oriented from left to
// right) and the vertical axis (oriented to the bottom)
func (c CoordinateSystem) canvasCoordinates(x, y float64) (px, py float64) {
	px = c.xorigin + c.xsign*c.unit2pixel*x
	py = c.yorigin + c.ysign*c.unit2pixel*y
	return px, py
}

func (c CoordinateSystem) canvasScaling(size float64) (psize float64) {
	psize = c.unit2pixel * size
	return psize
}

func (c CoordinateSystem) userCoordinates(px, py float64) (x, y float64) {
	x = (px - c.xorigin) / (c.xsign * c.unit2pixel)
	y = (py - c.yorigin) / (c.ysign * c.unit2pixel)
	return x, y
}

// UserCoordinatesBoundaries returns the boundaries of the (cnvwidth x
// cnvheight) canvas in user coordinates system: xmin, xmax, ymin, ymax. These
// values depends on 1/ the size of the canvas and 2/ the user coordinates
// system (placement of the origin, and length unit).
//
// It is be computed by retrieving the position of the top left corner and the
// bottom rigth corner in the user coordinates system
func (c CoordinateSystem) UserCoordinatesBoundaries() (xmin, xmax, ymin, ymax float64) {
	c1x, c1y := c.userCoordinates(0, 0)
	c2x, c2y := c.userCoordinates(float64(c.cnvxsize), float64(c.cnvysize))

	if c1x < c2x {
		xmin = c1x
		xmax = c2x
	} else {
		xmin = c2x
		xmax = c1x
	}

	if c1y < c2y {
		ymin = c1y
		ymax = c2y
	} else {
		ymin = c2y
		ymax = c1y
	}

	return xmin, xmax, ymin, ymax
}

/*
   def xyboundaries(self):
       """Returns the boundaries of the (cnvwidth x cnvheight) canvas in
       user coordinates system: xmin, xmax, ymin, ymax. These values
       depends on 1/ the size of the canvas and 2/ the user coordinates
       system (placement of the origin, and length unit).

       It is be computed by retrieving the position of the top left
       corner and the bottom rigth corner in the user coordinates
       system"""
       c1x, c1y = self.coordinatesSystem.xyCoordinates(0,0)
       c2x, c2y = self.coordinatesSystem.xyCoordinates(self.cnvwidth, self.cnvheight)

       if c1x < c2x:
           xmin = c1x
           xmax = c2x
       else:
           xmin = c2x
           xmax = c1x

       if c1y < c2y:
           ymin = c1y
           ymax = c2y
       else:
           ymin = c2y
           ymax = c1y

       return xmin, xmax, ymin, ymax
*/

// o2s return the sign to consider for the user axis orientation depending if
// the user axis is oriented as the canvas native axis (inverse = false) or
// inverse of the canvas native axis (inverse = true). If inverse is true, then
// it returns -1, else it return +1
func o2s(inverse bool) (sign float64) {
	if inverse {
		sign = -1
	} else {
		sign = +1
	}
	return sign
}

func newCoordSysAtOrigin(cnvXorigin, cnvYorigin float64, cnvwidth, cnvheight int, xrange float64) *CoordinateSystem {
	var xinverse bool = false // oriented as the canvas native orientation
	var yinverse bool = true  // inverse of the canvas native orientation
	xsign := o2s(xinverse)
	ysign := o2s(yinverse)

	// we consider a user x axis form 0 to xrange, then the unit has a size in
	// pixels equal to the canvas width divided by the xrange
	unit2pixel := float64(cnvwidth) / xrange
	// As a consequence, x and y are supposed to be values between 0 and xrange,
	// otherwize, they will not be viewed inside the canvas boundaries

	return &CoordinateSystem{
		cnvxsize:   cnvwidth,
		cnvysize:   cnvheight,
		xorigin:    cnvXorigin,
		yorigin:    cnvYorigin,
		xsign:      xsign,
		ysign:      ysign,
		unit2pixel: unit2pixel,
	}
}

// NewCoordSysBottomLeft creates a Coordinate System whose origin is at the
// bottom left corner of the canvas, with y coordinates axis oriented bottom up
// (inverse of the canvas native Y axis). The xrange is the range of x values
// (xmax - xmin) from the left boundary of the canvas to the right boundary.
func NewCoordSysBottomLeft(cnvwidth, cnvheight int, xrange float64) *CoordinateSystem {
	cnvXorigin := 0.
	cnvYorigin := float64(cnvheight)
	return newCoordSysAtOrigin(cnvXorigin, cnvYorigin, cnvwidth, cnvheight, xrange)
}

// NewCoordSysCentered creates a Coordinate System whose origin is at the center
// point of the canvas, with y coordinates axis oriented bottom up (inverse of
// the canvas native Y axis). The xrange is the range of x values (xmax - xmin)
// from the left boundary of the canvas to the right boundary.
func NewCoordSysCentered(cnvwidth, cnvheight int, xrange float64) *CoordinateSystem {
	cnvXorigin := float64(cnvwidth) * 0.5
	cnvYorigin := float64(cnvheight) * 0.5
	return newCoordSysAtOrigin(cnvXorigin, cnvYorigin, cnvwidth, cnvheight, xrange)
}

func NewCoordSysWithRanges(cnvwidth int, xmin, ymin, xmax, ymax float64) *CoordinateSystem {
	// See the test TestCoordSys_ImplementationOfWithRanges for explanation of the
	// two implementation, and the reason why we choose the second (the first
	// should be considered as deprecated)
	//return newCoordSysWithRanges_impl01(cnvwidth, xmin, ymin, xmax, ymax)
	return newCoordSysWithRanges_impl02(cnvwidth, xmin, ymin, xmax, ymax)
}

func newCoordSysWithRanges_impl01(cnvwidth int, xmin, ymin, xmax, ymax float64) *CoordinateSystem {
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

	return newCoordSysAtOrigin(cnvXorigin, cnvYorigin, cnvwidth, cnvheight, xrange)
}

func newCoordSysWithRanges_impl02(cnvwidth int, xmin, ymin, xmax, ymax float64) *CoordinateSystem {
	xrange := (xmax - xmin)
	yrange := (ymax - ymin)
	cnvheight := int((yrange / xrange) * float64(cnvwidth))

	// Create first a Coorinates System with origin at (0,0), so that we can use
	// the internal scaling functions of the coordinate system to reset
	// correctly the origin
	cs := newCoordSysAtOrigin(0, 0, cnvwidth, cnvheight, xrange)

	// Compute the pixel position of the point xmin,ymin in this coordinate system
	pxmin, pymin := cs.canvasCoordinates(xmin, ymin)

	// We can then determine the real origin pixel position with:
	pxOrigin := 0.
	pyOrigin := float64(cnvheight)
	cnvXorigin := pxOrigin - pxmin
	cnvYorigin := pyOrigin - pymin

	// And finally reset the origin of the coordinate system to the good value
	cs.xorigin = cnvXorigin
	cs.yorigin = cnvYorigin

	return cs
}

func NewCoordSysBoundedBy(cnvwidth int, points []struct{ X, Y float64 }, xoffset, yoffset float64) *CoordinateSystem {
	xmin, ymin, xmax, ymax := boundingBox(points)
	xmin = xmin - xoffset
	xmax = xmax + xoffset
	ymin = ymin - yoffset
	ymax = ymax + yoffset
	return NewCoordSysWithRanges(cnvwidth, xmin, ymin, xmax, ymax)
}

func boundingBox(points []struct{ X, Y float64 }) (xmin, ymin, xmax, ymax float64) {
	xmin = math.Inf(+1)
	xmax = math.Inf(-1)
	ymin = math.Inf(+1)
	ymax = math.Inf(-1)
	for _, p := range points {
		if p.X < xmin {
			xmin = p.X
		}
		if p.X > xmax {
			xmax = p.X
		}
		if p.Y < ymin {
			ymin = p.Y
		}
		if p.Y > ymax {
			ymax = p.Y
		}
	}
	return xmin, ymin, xmax, ymax
}
