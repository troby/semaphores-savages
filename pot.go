package main

type Pot struct {
	Servings	int
	Ready		chan bool
	Fill		chan bool
}

func (p *Pot) Init() {
	p.Servings	= 0
	p.Ready		= make(chan bool)
	p.Fill		= make(chan bool)
}
