package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import (
	"math"
)

type SidesType uint8

const (
	SidesTriangle SidesType = 3
	SidesSquare   SidesType = 4
	SidesCircle   SidesType = 0
)

func CalcSquare(sideLen float64, sidesNum SidesType) float64 {
	switch sidesNum {
	case 0:
		return 2 * math.Pi * math.Pow(sideLen, 2)
	case 3:
		return 0.5 * sideLen * sideLen * math.Sin(60.0)
	case 4:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
