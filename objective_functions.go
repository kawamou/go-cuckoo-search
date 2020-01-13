package main

// ObjectiveFunc is xxx
func ObjectiveFunc(name string) func([]float64) float64 {
	switch name {
	case "Ackley":
		return ackley
	default:
		return ackley
	}
}

func ackley(X []float64) float64 {
	var s1 float64
	for _, x := range X {
		s1 += x
	}
	return s1
}