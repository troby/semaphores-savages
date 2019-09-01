package main

import (
	"log"
	"sync"
	"time"
)

type Savage struct {
	Id		int
	WaitGroup	*sync.WaitGroup
	Write		*sync.Mutex
	Turn		*sync.Mutex
	Log		*log.Logger
	Pot		*Pot
}

func (s *Savage) Init(id int, wg *sync.WaitGroup, write *sync.Mutex, turn *sync.Mutex, logger *log.Logger, pot *Pot) {
	s.Id		= id
	s.WaitGroup	= wg
	s.Write		= write
	s.Turn		= turn
	s.Log		= logger
	s.Pot		= pot
}

func (s *Savage) GetFood() {
	defer s.WaitGroup.Done()
	s.Turn.Lock()
	if s.Pot.Servings == 0 {
		s.Write.Lock()
		s.Log.Print("Waking up the cook.")
		s.Write.Unlock()
		s.Pot.Fill <- true
		<-s.Pot.Ready
	}
	s.Write.Lock()
	s.Log.Printf("Savage %d is filling their bowl.\n", s.Id)
	s.Write.Unlock()
	s.Pot.Servings--
	s.Turn.Unlock()
	s.Eat()
}

func (s *Savage) Eat() {
	s.Write.Lock()
	s.Log.Printf("Savage %d is eating. There are %d servings remaining.\n", s.Id, s.Pot.Servings)
	s.Write.Unlock()
	time.Sleep(1 * time.Second)
}
