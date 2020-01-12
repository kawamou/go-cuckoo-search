package cs

import (
	"fmt"
)

func main() {
	objectiveFunc := "Rastrigin"
	solver := newSolver()
	solver.Run(objectiveFunc)
	fmt.Println(solver.getBest)
}