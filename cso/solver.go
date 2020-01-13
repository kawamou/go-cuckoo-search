package main

import (
	"math/rand"
	"time"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// TargetFunc is xxx
type TargetFunc func([]float64) float64

type swarm []*cuckoo

// Solver is xxx
type Solver struct {
	Ncuckoo     int         `yaml:"Ncuckoo"`
	Nstep       int         `yaml:"Nstep"`
	Ndim        int         `yaml:"Ndim"`
	bestFitness float64
	CuckooParam CuckooParam `yaml:"CuckooParam"`
	swarm       swarm
	targetFunc  TargetFunc
	UpperLimit  float64     `yaml:"UpperLimit"`
	LowerLimit  float64     `yaml:"LowerLimit"`
}

// CuckooParam is xxx
type CuckooParam struct {
	Stepsize float64 `yaml:"Stepsize"`
	Beta     float64 `yaml:"Beta"`
	Pa       float64 `yaml:"Pa"`
	SigmaP   float64 
	SigmaQ   float64 
}

func randomVector(nDim int, upperLimit, lowerLimit float64) []float64 {
	rand.Seed(time.Now().UnixNano())
	vector := make([]float64, nDim)
	for i := range vector {
		vector[i] = lowerLimit + rand.Float64()*(upperLimit-lowerLimit)
	}
	return vector
}

func(s *Solver) initSwarm() {
	var min float64
	swarm := make(swarm, s.Ncuckoo)

	for i := range swarm {
		pos := randomVector(s.Ndim, s.UpperLimit, s.LowerLimit)
		fitness := s.targetFunc(pos)
		swarm[i] = NewCuckoo(pos, fitness, s.CuckooParam)
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
func(s *Solver) step() {
	for i := range s.swarm {
		fitness := s.swarm[i].move(s.targetFunc)
		if fitness < s.bestFitness {
			s.bestFitness = fitness
		}
		fitness = s.swarm[i].randomMove(s.targetFunc, s.bestFitness)
		if fitness < s.bestFitness {
			s.bestFitness = fitness
		}
	}
}

// Run is xxx
func(s *Solver) Run() {
	s.initSwarm()
	for i := 0; i < s.Nstep; i++ {
		s.step()
		fmt.Println(s.bestFitness)
	}
}

// NewSolver is xxx
func NewSolver(targetFunc TargetFunc) *Solver {
	filepath := "./configs/config.yml"
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	c := new(Solver)
	err = yaml.Unmarshal(buf, &c)
	c.targetFunc = targetFunc
	fmt.Println(c)
	return c
}