// Example code filo_seq.go - Sequential Dinning Philosophers
// Since is sequential, elapsed timne could be simply calculated by (roubnds*philosophers*(time_to_think+time_to_eat))
// By Cesar De Rose

package main

import (
	"fmt"
	"time"
)

// Tempo de espera é 0 e só um filósofo come por vez
func Philosopher(id, iteration int) {
	time_to_think := 2
	time_to_eat := 3

	// think
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteration) // Thinking
	time.Sleep(time.Duration(time_to_think) * time.Second)

	//medição de tempo
	inicio := time.Now()

	// eat
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	eatingNow ++
	espera := time.Since(inicio)
	fmt.Printf(" E%d\n", iteration) // Eating
	ciclesPhilosophers[id] ++
	time.Sleep(time.Duration(time_to_eat) * time.Second)
}

func main() {
	philosophers := 5
	rounds := 20 //20 (cada filósofo executar 20 ciclos)
	ciclesPhilosophers int := make([]int, philosophers)
	mean int
	maxEating int
	longestWait int
	eatingNow int


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
		printf(cilesPhilosophers[i])
	}

	printf(mean)
	printf(maxEating)
	printf(longestWait)

	fmt.Printf("\nSequential Dinner took %s\n\n", elapsed)
}
