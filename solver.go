package cs

import (
	"log"
	"math/rand"
	"time"
	"io/ioutil"
	"path/filepath"
	"gopkg.in/yaml.v2"
)

// Solver is xxx
type Solver struct {
	Ncuckoo    int        `yaml: Ncuckoo`
	Nstep      int        `yaml: Nstep`
	Ndim       int        `yaml: Ndim`
	bestIndex  int
	Parameters []float64  `yaml: Parameters`
	swarm      swarm
	objectiveFunc func([]float64) float64
}

type swarm []*cuckoo

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
	var bestIndex int
	swarm := make(swarm, s.Ndim)

	for i, cuckoo := range s.swarm {
		pos := RandomVector(s.nDim, 0, 1)
		fitness := s.objectiveFunc(pos)
		cuckoo := NewCuckoo(pos, fitness, s.Parameters[0], s.Parameters[1])
		if i := 0 {
			min = fitness
			bestIndex = i
		} else {
			if fitness < min {
				min = fitness
				bestIndex = i
			}
		}
	}
	s.swarm = swarm
	s.bestIndex = bestIndex
}

// Step is xxx
func(s *Solver) Step() {
	for cuckoo := range s.swarm {
		cuckoo.Step()
	}
}

func(s *Solver) evalFitness() {

}

func(s *Solver) randomMove() {

}

// Minimize is xxx
func(s *Solver) Run() {
	s.initSwarm()
	for i := 0; i < s.Nstep; i++ {
		s.step()
		s.randomMove()
		log.Print()
	}
}

// constructor
func newsolver() *Solver {
	filepath := "./configs/config.yml"
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	c := new(Solver)
	err := yaml.Unmarshal(buf, &c)
	return c
}