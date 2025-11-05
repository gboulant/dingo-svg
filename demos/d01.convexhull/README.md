# Computation (and visualisation) of a convex hull

This programs illustrates an usage of the svg sketcher to control the result of
a geometric algorithm. The algorithm is the computation of the convex hull of a
set of points of a 2D plane see the implementation in the file
[algorithm.go](algorithm.go).

Even if the main objective is to illustrate the functions of the svg sketcher,
this example shows how to implement some basic geometric operations. Indeed, the
convex hull algorithms requires vectorial calculus like scalar product, vector
product, computing the angle between two vectors, etc (see the file
[vector.go](vector.go)).