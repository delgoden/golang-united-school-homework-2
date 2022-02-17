package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type numSides uint

const (
	SidesTriangle numSides = 3
	SidesSquare   numSides = 4
	SidesCircle   numSides = 0
)

func CalcSquare(sideLen float64, sidesNum numSides) (square float64) {
	switch sidesNum {
	case 3:
		square = math.Pow(3.0, 0.5) / 4.0 * math.Pow(sideLen, 2)
	case 4:
		square = math.Pow(sideLen, 2)
	case 0:
		square = math.Pi * math.Pow(sideLen, 2)
	default:
		square = 0
	}
	return
}
