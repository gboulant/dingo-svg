# Using go runtime profiling tool

The creation of the 3D surface sketch (demo01) is time consuming (computation on
a 100x100 grid). Then it is a good exercice to experiment the go pprof tools.
After analysis, the time consuming instruction appears to be the instruction
that cumulates the SVG polygon description lines in the string variable `body`
(see file `sketcher.go`):

```go
s.body += fmt.Sprintf(polygPattern, coords, style) + "\n"
```

A standard way for profiling an application is to use the package
`runtime/pprof`, together with the go tool pprof. A good starting point is to
read the reference documentation page at https://pkg.go.dev/runtime/pprof. You
may also read the following document that gives guide lines for using the go
pprof tool:

* https://blog.stackademic.com/profiling-go-applications-in-the-right-way-with-examples-e784526e9481

In short, when using go test and/or bench, the profiling is natively activable
using the options `-cpuprofile` and `-memprofile`:

```shell
go test -bench='.' -cpuprofile='cpu.prof' -memprofile='mem.prof'
```

For profiling a standalone program, you have to explicitly integrate the
`runtime/pprof` functions in your main function, as described in the reference
documentation https://pkg.go.dev/runtime/pprof.

For convenience (and also to make sure to have correctly understand), the
standard usage of the `runtime/pprof` functions is packaged into a `Profiler`
tool implemented in the file [profiling.go](profiling.go). At the begining of
your main function, just add the following code:

```go
func main() {
	p := NewProfiler("output.cpu.prof", "output.mem.prof")
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
	defer p.Stop()

	// The rest of your pogram ...
}
```

Then, when the program has ended, you have two files `output.cpu.prof` and
`output.mem.prof` that you can analyse using the go tool pprof. Some standard
commands hereafter.

Start the analysis of the CPU profiling:

```shell
go tool pprof output.cpu.prof
```

That executes an interactive program:

```shell
File: d02.isometry
Build ID: 40ee0bccf7830da33ef5f4731ec9344e2fdbe186
Type: cpu
Time: 2025-12-01 17:02:49 CET
Duration: 4.51s, Total samples = 8.14s (180.29%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) 
```

To get the top 15 cpu consuming functions, focusing on the functions of the
program (hiding the runtime functions):

```shell
(pprof) hide=runtime
(pprof) top15
```

That displays:

```shell
Active filters:
   hide=runtime
Showing nodes accounting for 3.85s, 47.30% of 8.14s total
Dropped 14 nodes (cum <= 0.04s)
Showing top 15 nodes out of 20
      flat  flat%   sum%        cum   cum%
     3.65s 44.84% 44.84%      3.84s 47.17%  github.com/gboulant/dingo-svg.(*Sketcher).Polygon
     0.05s  0.61% 45.45%      3.89s 47.79%  main.IsometricView.DrawPolygon
     0.04s  0.49% 45.95%      0.04s  0.49%  strconv.rightShift
     0.03s  0.37% 46.31%      0.15s  1.84%  fmt.(*pp).printArg
     0.03s  0.37% 46.68%      0.07s  0.86%  strconv.(*decimal).Shift
     0.02s  0.25% 46.93%      0.02s  0.25%  fmt.(*buffer).write (inline)
     0.01s  0.12% 47.05%      0.17s  2.09%  fmt.(*pp).doPrintf
     0.01s  0.12% 47.17%      0.04s  0.49%  github.com/gboulant/dingo-svg.Pencil.DrawStyleWithFillMode
     0.01s  0.12% 47.30%      0.09s  1.11%  strconv.bigFtoa
         0     0% 47.30%      0.11s  1.35%  fmt.(*fmt).fmtFloat
         0     0% 47.30%      0.02s  0.25%  fmt.(*fmt).pad
         0     0% 47.30%      0.11s  1.35%  fmt.(*pp).fmtFloat
         0     0% 47.30%      0.18s  2.21%  fmt.Sprintf
         0     0% 47.30%      3.90s 47.91%  main.DrawIsometricView
         0     0% 47.30%      0.92s 11.30%  main.demo01_cardinalsine
(pprof) 
```

We can then focus on the functions `Polygon` that are the time consuming
functions using the command `list`

```shell
(pprof)  list Polygon
Total: 8.14s
ROUTINE ======================== github.com/gboulant/dingo-svg.(*Sketcher).Polygon in /home/guillaume/develspace/wkspace.dingo/go-packages/svg/sketcher.go
     3.65s      3.84s (flat, cum) 47.17% of Total
         .          .    189:func (s *Sketcher) Polygon(points []struct{ X, Y float64 }, fill bool) {
         .          .    190:   var x, y float64
         .          .    191:   var coords string
         .          .    192:   for _, p := range points {
         .          .    193:           x, y = p.X, p.Y
         .          .    194:           px, py := s.canvasCoordinates(x, y)
      50ms      190ms    195:           coords += fmt.Sprintf("%.2f,%.2f ", px, py)
         .          .    196:   }
      10ms       50ms    197:   style := s.Pencil.DrawStyleWithFillMode(fill)
     3.59s      3.60s    198:   s.body += fmt.Sprintf(polygPattern, coords, style) + "\n"
         .          .    199:   s.x = x
         .          .    200:   s.y = y
         .          .    201:}
         .          .    202:
         .          .    203:// Polyline draws a continuous line made of multiple connected edges,
```
