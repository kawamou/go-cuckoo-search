package main

import (
	"go-cuckoo-search/cso"
)

func main() {
	targetFunc := cso.GetTargetFunc("Ackley")
	solver := cso.NewSolver(targetFunc)
	solver.Run()
}