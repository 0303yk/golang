package rectangle

import "math"

func Area(wid ,len float64) float64 {
	area := wid * len
	return area
}

func Diagonal(wid ,len float64) float64 {
	diagonal := math.Sqrt((wid * wid)+ (len*len))
	return diagonal
}