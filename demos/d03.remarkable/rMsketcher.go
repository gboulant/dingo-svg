package main

import svg "github.com/gboulant/dingo-svg"

const (
	RM_WIDTH_PX         = 1404 // width of remarkable templates
	RM_HEIGHT_PX        = 1872 // height or remarkable templates
	RM_LEFTBAR_WIDTH_PX = 120  // width of the left bar

	RM_RESOLUTION_DPI    = 226 // DPI (digit per inch, i.e. pixel per inch)
	MILLIMETERS_PER_INCH = 25.4
)

func pixel2inch(pixel float64) float64 {
	return pixel / RM_RESOLUTION_DPI
}

func inch2millimeter(inch float64) float64 {
	return inch * MILLIMETERS_PER_INCH
}

var (
	RM_WIDTH_INCH       float64 = pixel2inch(RM_WIDTH_PX)
	RM_WIDTH_MILLIMETER float64 = inch2millimeter(RM_WIDTH_INCH)
)

func NewRemarkableSketcher() *svg.Sketcher {
	cs := svg.NewCoordSysBottomLeft(RM_WIDTH_PX, RM_HEIGHT_PX, RM_WIDTH_MILLIMETER)
	sk := svg.NewSketcher().WithCoordinateSystem(cs)
	sk.WithBackgroundColor("white")
	return sk
}
