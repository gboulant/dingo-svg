package main

import "log"

func main() {
	var demo func() error

	demo = demo_rM_millimeters

	if err := demo(); err != nil {
		log.Fatal(err)
	}
}
