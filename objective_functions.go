package main

import (
	"math"
)

// GetTargetFunc is xxx
func GetTargetFunc(name string) func([]float64) float64 {
	switch name {
	case "Ackley":
		return ackley
	default:
		return ackley
	}
}

func ackley(X []float64) float64 {
    a, b, c, d := 20.0, 0.2, 2*math.Pi, float64(len(X))
    var s1, s2 float64
    for _, x := range X {
        s1 += x * x
        s2 += math.Cos(c * x)
    }
    return -a*math.Exp(-b*math.Sqrt(s1/d)) - math.Exp(s2/d) + a + math.Exp(1)
}