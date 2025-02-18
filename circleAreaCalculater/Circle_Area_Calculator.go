package circleareacalculater

import "math"

func CalculateArea(radius float64) float64 {
	const PI float64 = math.Pi;

	return PI * radius * radius;
}