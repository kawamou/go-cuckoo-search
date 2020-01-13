package main

import (
	"fmt"
	"math/rand"
	"time"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// Solver is xxx
type Solver struct {
	Ncuckoo    int        `yaml:"Ncuckoo"`
	Nstep      int        `yaml:"Nstep"`
	Ndim       int        `yaml:"Ndim"`
	bestFitness float64
	Parameters []float64  `yaml:"Parameters"`
	swarm      swarm
	objectiveFunc func([]float64) float64
}

type swarm []*cuckoo

// RandomVector is xxx
func RandomVector(nDim int, upperLimit, lowerLimit float64) []float64 {
	rand.Seed(time.Now().UnixNano())
	vector := make([]float64, nDim)
	for i := range vector {
		vector[i] = lowerLimit + rand.Float64()*(upperLimit-lowerLimit)
	}
	return vector
}

func(s *Solver) initSwarm() {
	var min float64
	swarm := make(swarm, s.Ndim)

	for i := range swarm {
		pos := RandomVector(s.Ndim, 0, 1)
		fitness := s.objectiveFunc(pos)
		swarm[i] = NewCuckoo(pos, fitness, s.Parameters[0], s.Parameters[1], s.Parameters[2])
		if i == 0 {
			min = fitness
		} else {
			if fitness < min {
				min = fitness
			}
		}
	}
	s.swarm = swarm
	s.bestFitness =  min
}

// Step is xxx
func(s *Solver) Step() {
	for i := range s.swarm {
		s.swarm[i].move()
		s.swarm[i].randomMove(s.bestFitness)
	}
}

func(s *Solver) evalFitness() {

}

// Run is xxx
func(s *Solver) Run() {
	s.initSwarm()
	for i := 0; i < s.Nstep; i++ {
		s.Step()
	}
}

// NewSolver is xxx
func NewSolver(objectiveFunc func([]float64) float64) *Solver {
	filepath := "./configs/config.yml"
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	c := new(Solver)
	fmt.Println(c)
	err = yaml.Unmarshal(buf, &c)
	fmt.Println(c)
	c.objectiveFunc = objectiveFunc
	return c
}