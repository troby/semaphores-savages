package main

import (
	"log"
	"os"
	"sync"
)

const (
	MAX_SERVINGS = 10
	MAX_SAVAGES = 32
)

func main() {
	wg := new(sync.WaitGroup)
	write := new(sync.Mutex)
	logger := log.New(os.Stderr, "Log: ", 0)
	turn := new(sync.Mutex)
	pot := new(Pot)
	cook := new(Cook)

	cook.Init(write, logger, pot)
	go cook.Work()

	wg.Add(MAX_SAVAGES)
	for count:=1; count<=MAX_SAVAGES; count++ {
		savage := new(Savage)
		savage.Init(count, wg, write, turn, logger, pot)
		go savage.GetFood()
	}

	wg.Wait()
	cook.Stop <- true
	<-cook.Done
}
