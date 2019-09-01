package main

import (
	"log"
	"sync"
)

type Cook struct {
	Write		*sync.Mutex
	Log		*log.Logger
	Pot		*Pot
	Stop		chan bool
	Done		chan bool
}

func (c *Cook) Init(write *sync.Mutex, logger *log.Logger, pot *Pot) {
	pot.Init()
	c.Write		= write
	c.Log		= logger
	c.Pot		= pot
	c.Stop		= make(chan bool)
	c.Done		= make(chan bool)
}

func (c *Cook) Work() {
	for {
		select {
		case <-c.Stop:
			c.Write.Lock()
			c.Log.Print("Stop signal received by cook.")
			c.Write.Unlock()
			c.Done <- true
			return
		case <-c.Pot.Fill:
			c.Write.Lock()
			c.Log.Print("The cook is filling the pot.")
			c.Write.Unlock()
			c.Pot.Servings = MAX_SERVINGS
			c.Pot.Ready <- true
		}
	}
}
