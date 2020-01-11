package cs

type cuckoo struct {
	position []float64
	fitness  []float64
}

func(cuckoo *cuckoo) updatePosition() {

}

func(cuckoo *cuckoo) updateFitness() {

}

func(cuckoo *cuckoo) setPosition() {

}

func(cuckoo *cuckoo) setFitness() {

}

func newCuckoo(position float64, fitness float64) *cuckoo {
	c := new(cuckoo{position, fitness})
	c.postition := position
	c.fitness := fitness
	return c
}