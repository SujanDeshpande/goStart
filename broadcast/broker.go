package main

type Broker struct {
	stopCh    chan interface{}
	publishCh chan interface{}
	subCh     chan interface{}
	unsubCh   chan interface{}
}

func NewBroker() *Broker {
	broker := &Broker{
		stopCh:    make(chan interface{}),
		publishCh: make(chan interface{}),
		subCh:     make(chan interface{}),
		unsubCh:   make(chan interface{}),
	}
	return broker
}
