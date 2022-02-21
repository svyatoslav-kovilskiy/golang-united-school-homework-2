package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
const (
	SidesTriangle = 3

	SidesSquare = 4

	SidesCircle = 0
)

func CalcTriangleSquare(sideLen float64) float64 {
	// S = (3)**0,5 * a**2 / 4
	S := math.Pow(3, 0.5) * math.Pow(sideLen, 2) / 4
	return S
}

func CalcRectangleSquare(sideLen float64) float64 {
	S := math.Pow(sideLen, 2)
	return S
}

func CalcCircleSquare(radius float64) float64 {
	S := math.Pi * math.Pow(radius, 2)
	return S
}

func CalcSquare(sideLen float64, sidesNum int) float64 {
	if sidesNum == 4 {
		return CalcRectangleSquare(sideLen)
	} else if sidesNum == 3 {
		return CalcTriangleSquare(sideLen)
	} else if sidesNum == 0 {
		return CalcCircleSquare(sideLen)
	} else {
		return 0.0
	}

}
