// Example code filo_seq.go - Sequential Dinning numFilosofos
// Since is sequential, elapsed timne could be simply calculated by (roubnds*numFilosofos*(tempoPensar+tempoComer))
// By Cesar De Rose

package main

import (
	"fmt"
	"time"
)

const (
	numFilosofos = 5
	numCiclos    = 20
	tempoPensar  = 100 * time.Millisecond
	tempoComer   = 150 * time.Millisecond
)

var ciclesnumFilosofos = make([]int, numFilosofos) //número de ciclos concluídos por cada filósofo
var eatingNow int
var somaEspera time.Duration //quanto foi esperado ao todo
var longestWait time.Duration
var maxEating int

// Tempo de espera é 0 e só um filósofo come por vez
// A “espera” da V0 vai sair como algo na casa de nanossegundos (por exemplo, 100ns),
// e não exatamente 0. É só o custo de chamar time.Now() e time.Since(), e é esperado.
func Filosofo(id, iteration int) {

	// think
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteration) // Thinking
	time.Sleep(tempoPensar)

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

	fmt.Printf(" E%d\n", iteration) // Eating

	time.Sleep(tempoComer)

	//aqui será a liberação dos garfos
	eatingNow--

	ciclesnumFilosofos[id]++
}

func main() {

	fmt.Printf("\n[P1] [P2] [P3] [P4] [P5]\n")

	// run numFilosofos

	start := time.Now()

	for i := 1; i <= numCiclos; i++ {
		for j := 0; j < numFilosofos; j++ {
			Filosofo(j, i)
		}
	}
	elapsed := time.Since(start)

	media := somaEspera / (numFilosofos * numCiclos)

	fmt.Printf("---Tempo médio de espera para aquisição dos garfos: %s\n", media)
	fmt.Printf("---Número máximo de filósofos comendo simultaneamente: %d\n", maxEating)
	fmt.Printf("---Maior tempo de espera observado: %s\n", longestWait)

	for i := 0; i < len(ciclesnumFilosofos); i++ {
		fmt.Printf("---Filósofo %d: %d ciclos\n", i, ciclesnumFilosofos[i])
	}

	fmt.Printf("\nSequential Dinner took %s\n\n", elapsed)
}
