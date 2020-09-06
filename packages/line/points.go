package main

import "math"

//As we already know, create a package called main allows us create the main function.
//About visibility:
//When the first letter is uppercase, it's public under and outside packages.
//When the first letter is downcase, it's private, can be accessed only under it package.
//For example...
//Public - to declare something public, uppercase in first letter.
//private or _private - to declare something private, downcase or _ in first letter.
//Important to know:
//Visibility exists only under the package, doesn't matter for the archives.
//It's recommended that every public declarations have a comment to identify.
//You can't have multiple functions with the same name under the same package.
//The archive it's not really important, what matters is the package.

//Represents a point in a cartesian plan.
type Point struct {
	x float64
	y float64
}

func cathetus (a, b Point) (cx, cy float64){
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y- a.y)
	return
}

func Distance (a, b Point) float64{
	cx, cy := cathetus(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}