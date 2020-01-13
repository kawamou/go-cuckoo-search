package main

import (
	"time"
	"math/rand"
)

type cuckoo struct {
	position []float64
	fitness  float64
	stepsize float64
	beta     float64
	pa       float64
}

func(c *cuckoo) randomWalk(position []float64) []float64 {
	return position
}

func(c *cuckoo) move(targetFunc TargetFunc) {
	oldPosition := c.position
	oldFitness := c.fitness
	newPosition := c.randomWalk(oldPosition)
	newFitness := targetFunc(newPosition)
	
	if newFitness < oldFitness {
		c.position = newPosition
		c.fitness = newFitness
	}
}

func(c *cuckoo) randomMove(targetFunc TargetFunc, bestFitness float64) {
	if c.fitness <= bestFitness {
		
	} else {
		rand.Seed(time.Now().UnixNano())
		if rand.Float64() < c.pa {
			c.position = c.randomWalk(c.position)
			c.fitness = targetFunc(c.position)
		}
	}
}

// NewCuckoo is constructor
func NewCuckoo(position []float64, fitness, stepsize, beta, pa float64) *cuckoo {
	c := &cuckoo{position, fitness, stepsize, beta, pa}
	return c
}