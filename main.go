package main

import (
	"fmt"
)

func main() {
	targetFunc := GetTargetFunc("Ackley")
	solver := NewSolver(targetFunc)
	solver.Run()
	fmt.Println(solver.bestFitness)
}