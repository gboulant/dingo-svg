package main

/*
Tool that implements the standard way to make runtime profiling of a go program
(using the standard package runtime/pprof). See https://pkg.go.dev/runtime/pprof.
*/

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

type Profiler struct {
	cpupath string
	mempath string
	cpufile *os.File
}

func NewProfiler(cpupath, mempath string) *Profiler {
	return &Profiler{cpupath, mempath, nil}
}

func (p *Profiler) Start() error {
	if err := p.StartCPUProfile(); err != nil {
		return err
	}
	// The MEM profile is done at the end of the process, using the function
	// CollectMEMProfiling
	return nil
}

func (p *Profiler) Stop() error {
	if err := p.CollectMEMProfile(); err != nil {
		return fmt.Errorf("error when collecting memory profiling (%s)", err)
	}
	if err := p.StopCPUProfile(); err != nil {
		return fmt.Errorf("error when stopping CPU profiling (%s)", err)
	}
	return nil
}

func (p *Profiler) StartCPUProfile() error {
	f, err := os.Create(p.cpupath)
	if err != nil {
		return fmt.Errorf("could not create CPU profile: %s", err)
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		return fmt.Errorf("could not start CPU profiling: %s", err)
	}

	p.cpufile = f
	return nil
}

func (p *Profiler) StopCPUProfile() error {
	pprof.StopCPUProfile()
	p.cpufile.Close()
	p.cpufile = nil
	return nil
}

func (p *Profiler) CollectMEMProfile() error {
	f, err := os.Create(p.mempath)
	if err != nil {
		return fmt.Errorf("could not create MEM profile: %s", err)
	}
	defer f.Close()

	runtime.GC() // get up-to-date statistics
	// Lookup("allocs") creates a profile similar to go test -memprofile.
	// Alternatively, use Lookup("heap") for a profile
	// that has inuse_space as the default index.
	if err := pprof.Lookup("allocs").WriteTo(f, 0); err != nil {
		return fmt.Errorf("could not write memory profile: %s", err)
	}

	return nil
}
