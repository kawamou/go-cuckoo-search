package main

import (
	"fmt"
)

func main() {
	objectiveFunc := ObjectiveFunc("Ackley")
	solver := NewSolver(objectiveFunc)
	solver.Run()
	fmt.Println(solver.bestFitness)
}