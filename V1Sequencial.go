// Example code filo_seq.go - Sequential Dinning Philosophers
// Since is sequential, elapsed timne could be simply calculated by (roubnds*philosophers*(time_to_think+time_to_eat))
// By Cesar De Rose

package main

import (
	"fmt"
	"time"
)

var ciclesPhilosophers = make([]int, philosophers)
var eatingNow int
var time_to_think = 2
var time_to_eat = 3
var somaEspera time.Duration //quanto foi esperado ao todo
var longestWait time.Duration
var maxEating int
var philosophers = 5
var rounds = 20 //20 (cada filósofo executar 20 ciclos)

// Tempo de espera é 0 e só um filósofo come por vez
func Philosopher(id, iteration int) {

	// think
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteration) // Thinking
	time.Sleep(time.Duration(time_to_think) * time.Second)

	//medição de tempo
	inicio := time.Now()

	//aqui seria aquisição dos garfos

	espera := time.Since(inicio)
	somaEspera += espera

	if espera > longestWait {
		longestWait = espera
	}

	eatingNow++
	if eatingNow > maxEating {
		maxEating = eatingNow
	}

	// eat
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}

	eatingNow--

	fmt.Printf(" E%d\n", iteration) // Eating

	ciclesPhilosophers[id]++

	time.Sleep(time.Duration(time_to_eat) * time.Second)
}

func main() {
	var mean int

	fmt.Printf("\n[P1] [P2] [P3] [P4] [P5]\n")

	// run philosophers

	start := time.Now()

	for i := 1; i <= rounds; i++ {
		for j := 0; j < philosophers; j++ {
			Philosopher(j, i)
		}
	}
	elapsed := time.Since(start)

	mean = rounds * philosophers * (time_to_think + time_to_eat)

	for i := 0; i < len(ciclesPhilosophers); i++ {
		printf(ciclesPhilosophers[i])
	}

	printf(mean)
	printf(maxEating)
	printf(longestWait)

	fmt.Printf("\nSequential Dinner took %s\n\n", elapsed)
}
