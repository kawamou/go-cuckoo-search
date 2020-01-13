package main

import (
)

func main() {
	targetFunc := GetTargetFunc("Ackley")
	solver := NewSolver(targetFunc)
	solver.Run()
}