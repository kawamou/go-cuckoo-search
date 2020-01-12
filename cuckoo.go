package cs

import (
	"time"
	"math/rand"
)

type cuckoo struct {
	position []float64
	fitness  []float64
	stepsize float64
	beta     float64
	pa       float64
}

func(c *cuckoo) evalFitness(position []float64) []float64 {
	return c.objectiveFunc(position)
}

func(c *cuckoo) RandomWork() {
}

func(c *cuckoo) Step() {
	cuckoo.move()
	cuckoo.randomMove()
}

func(c *cuckoo) move() {
	oldPosition := c.position
	oldFitness := c.fitness
	newPosition := c.randomWork(oldPosition)
	newFitness := c.evalFitness(newPosition)
	
	if newFitness < oldFitness {
		c.position = newPosition
		c.fitness = newFitness
	}

	if newFitness < c.bestFitness {
		c.bestPosition = newPosition
		c.bestFitness = newFitness
	}
}

func(c *cuckoo) randomMove() {
	if c.fitness =< c.bestFitness {
		break
	} else {
		rand.Seed(time.Now().UnixNano())
		if rand.float64 < c.pa {
			c.position := c.randomWork(oldPosition)
			c.fitness := c.evalFitness(newPosition)
		}
	}
}

// NewCuckoo is constructor
func NewCuckoo(position, fitness []float64, stepsize, beta float64) *cuckoo {
	c := &cuckoo{position, fitness, stepsize, beta}
	return c
}