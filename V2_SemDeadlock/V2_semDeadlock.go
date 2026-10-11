package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numFilosofos = 5
	numCiclos    = 20
	tempoPensar = 100 * time.Millisecond
	tempoComer  = 150 * time.Millisecond
	
)

var garfos [numFilosofos]sync.Mutex
var mesa sync.Mutex
var somaEspera [numFilosofos]time.Duration
var maiorEspera [numFilosofos]time.Duration

func filosofo(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	esq := id
	dir := (id + 1) % numFilosofos

	for ciclo := 1; ciclo <= numCiclos; ciclo++ {
		// 1. pensar
		fmt.Printf("Filósofo %d está pensando\n", id)
		time.Sleep(tempoPensar)

		inicioEspera := time.Now()
		mesa.Lock()

		// 2. pegar o garfo esquerdo
		fmt.Printf("Filósofo %d quer o garfo %d\n", id, esq)
		garfos[esq].Lock()
		fmt.Printf("Filósofo %d pegou o garfo %d\n", id, esq)

		// 3. pegar o garfo direito
		fmt.Printf("Filósofo %d quer o garfo %d\n", id, dir)
		garfos[dir].Lock()
		fmt.Printf("Filósofo %d pegou garfo %d\n", id, dir)

		espera := time.Since(inicioEspera)
		somaEspera[id] += espera
		if espera > maiorEspera[id] {
			maiorEspera[id] = espera
		}

		// 4. comer
		fmt.Printf("Filósofo %d está comendo\n", id)
		time.Sleep(tempoComer)
	
		// 5. soltar os dois garfos
		garfos[esq].Unlock()
		fmt.Printf("Filósofo %d soltou o garfo %d\n", id, esq)
		garfos[dir].Unlock()
		fmt.Printf("Filósofo %d soltou o garfo %d\n", id, dir)

		mesa.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < numFilosofos; i++ {
		wg.Add(1)
		go filosofo(i, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Tempo: %s\n", elapsed)

	var total time.Duration
	var maior time.Duration

	for i := 0; i < numFilosofos; i++ {
		total += somaEspera[i]
		if maiorEspera[i] > maior {
			maior = maiorEspera[i]
		}
	}

	media := total / time.Duration(numFilosofos*numCiclos)
	fmt.Printf("Espera média: %s\n", media)
	fmt.Printf("Maior espera: %s\n", maior)

}