package main

import (
	"fmt"
)

type catalog []struct {
	name string
	make func(svgpath string) error
}

func main() {
	var rMsketchers catalog = catalog{
		{"rM_millimeters", rM_millimeters},
		{"rM_ecolier", rM_ecolier},
	}

	for _, sketcher := range rMsketchers {
		svgpath := fmt.Sprintf("output.%s.svg", sketcher.name)
		if err := sketcher.make(svgpath); err != nil {
			fmt.Printf("err: sketcher %s failed du to error %s\n", sketcher.name, err)
		}
	}
}
