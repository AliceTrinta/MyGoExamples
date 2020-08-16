package main

import "math"

//When the first letter is uppercase, it's public
//under and outside packages.
//You can't have multiple functions with the same name
//under the same package.
//The archive it's not really important, what matters is the package.
//When the first letter is downcase, it's private to only the package
//For example...
//Public - to declare something public, uppercase in first letter.
//private or _private - to declare something private, downcase or _ in first letter.
//Every public declarations have to have a comment to identify.

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