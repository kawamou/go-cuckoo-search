package main

import (
	// "fmt"
	"time"
	"math/rand"
	"math"
)

type cuckoo struct {
	position    []float64
	fitness     float64
	CuckooParam
}

func(c *cuckoo) levyFright(position float64) float64 {
	p := rand.NormFloat64() * c.SigmaP
	q := rand.NormFloat64() * c.SigmaQ
	levy := p / math.Pow(math.Abs(q), 1.0/c.Beta)
	return levy
}

func(c *cuckoo) randomWalk(position []float64) []float64 {
	for i := range position {
		position[i] += c.Stepsize * c.levyFright(position[i])
	}
	return position
}

func(c *cuckoo) move(targetFunc TargetFunc) float64 {
	oldPosition := c.position
	oldFitness := c.fitness
	newPosition := c.randomWalk(oldPosition)
	newFitness := targetFunc(newPosition)
	
	if newFitness < oldFitness {
		c.position = newPosition
		c.fitness = newFitness
		// fmt.Println(c.fitness)
	}

	return c.fitness
}

func(c *cuckoo) randomMove(targetFunc TargetFunc, bestFitness float64) float64 {
	if c.fitness <= bestFitness {
		return c.fitness
	} else {
		rand.Seed(time.Now().UnixNano())
		if rand.Float64() < c.Pa {
			c.position = c.randomWalk(c.position)
			c.fitness = targetFunc(c.position)
		}
		return c.fitness
	}
}

// NewCuckoo is constructor
func NewCuckoo(position []float64, fitness float64, cuckooParam CuckooParam) *cuckoo {
	c := &cuckoo {
		position:    position,
		fitness:     fitness,
		CuckooParam: CuckooParam {
			Stepsize: cuckooParam.Stepsize,
			Beta: cuckooParam.Beta,
			Pa: cuckooParam.Pa,
			SigmaP: cuckooParam.SigmaP,
			SigmaQ: cuckooParam.SigmaQ,
		}}
	Beta := c.Beta
	numerator := math.Gamma(1.0+Beta)*math.Sin(math.Pi*Beta/2.0)
	denominator := math.Pow(math.Gamma((1.0+Beta)/2.0)*Beta*2.0, (Beta-1.0)/2.0)
	c.SigmaP = math.Pow(numerator / denominator , 1.0/Beta)
	c.SigmaQ = 1.0
	return c
}