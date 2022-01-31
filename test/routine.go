package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Start ")
	pingCh := make(chan bool, 1)
	pongCh := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go ping(pingCh, pongCh, &wg)
	go pong(pingCh, pongCh, &wg)
	pongCh <- true
	wg.Wait()
	fmt.Println("Complete ")
}

func ping(pingCh chan bool, pongCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range pingCh {
		fmt.Println("Ping", v)
		pongCh <- true
	}
}

func pong(pingCh chan bool, pongCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range pongCh {
		fmt.Println("Pong", v)
		pingCh <- true
	}
}
