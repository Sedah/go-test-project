package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(CalArea(10.0, 30.0))
}

func CalArea(radius, angle float64) string {
	r := AreaOfSegment(radius, angle)
	r2 := AreaOfSegment(radius, (360.0 - angle))
	return fmt.Sprintf("%.4f, %.4f", r, r2)
}

func AreaOfSegment(radius, angle float64) float64 {
	pi := 3.14159
	//Calculating area of sector
	area_of_sector := pi * (radius * radius) * (angle / 360.0)

	//Calculating area of triangle
	area_of_triangle := 1.0 / 2.0 * (radius * radius) * math.Sin((angle*pi)/180.0)
	return area_of_sector - area_of_triangle
}
