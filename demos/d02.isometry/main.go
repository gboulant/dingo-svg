package main

import "log"

func main() {
	p := NewProfiler("output.cpu.prof", "output.mem.prof")
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
	defer p.Stop()

	demo01_cardinalsine()
	demo01_horseshoe()
	demo01_parabol()
	demo02()
}
