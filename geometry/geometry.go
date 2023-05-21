package geometry

import "math"

// Area is exported because it starts with a capital letter
func Area(length, width float64) float64 {
	return length * width
}

// Perimeter is exported because it starts with a capital letter
func Perimeter(length, width float64) float64 {
	return 2 * (length + width)
}

// Diagonal is exported because it starts with a capital letter
func Diagonal(length, width float64) float64 {
	return math.Sqrt((length * length) + (width * width))
}

// init is called before main
func init() {
	println("geometry package initialized")
}
