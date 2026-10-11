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
	tempoPausa  = 50 * time.Millisecond
)

var garfos [numFilosofos]sync.Mutex

func filosofo(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	esq := id
	dir := (id + 1) % numFilosofos

	for ciclo := 1; ciclo <= numCiclos; ciclo++ {
		// 1. pensar
		fmt.Printf("Filósofo %d está pensando\n", id)
		time.Sleep(tempoPensar)

		// 2. pegar o garfo esquerdo
		fmt.Printf("Filósofo %d quer o garfo %d\n", id, esq)
		garfos[esq].Lock()
		fmt.Printf("Filósofo %d pegou o garfo %d\n", id, esq)

		// 3. pausa curta (para provocar o deadlock)
		fmt.Printf("Filósofo %d está parado\n", id)
		time.Sleep(tempoPausa)

		// 4. pegar o garfo direito
		fmt.Printf("Filósofo %d quer o garfo %d\n", id, dir)
		garfos[dir].Lock()
		fmt.Printf("Filósofo %d pegou garfo %d\n", id, dir)

		// 5. comer
		fmt.Printf("Filósofo %d está comendo\n", id)
		time.Sleep(tempoComer)
	
		// 6. soltar os dois garfos
		garfos[esq].Unlock()
		fmt.Printf("Filósofo %d soltou o garfo %d\n", id, esq)
		garfos[dir].Unlock()
		fmt.Printf("Filósofo %d soltou o garfo %d\n", id, dir)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < numFilosofos; i++ {
		wg.Add(1)
		go filosofo(i, &wg)
	}
	wg.Wait()
}