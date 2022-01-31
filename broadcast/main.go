package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Start")
	subCh := make(chan interface{})
	unsubCh := make(chan interface{})
	valueChs := make([]chan interface{}, 0)

	for i := 0; i < 10; i++ {
		valueCh := make(chan interface{})
		go worker(valueCh)
		valueChs = append(valueChs, valueCh)
	}
	go publish(subCh, unsubCh, valueChs)

	for i := 0; i < 10; i++ {
		s := "sujan " + strconv.Itoa(i)
		subCh <- s
		unsubCh <- s
	}

}

func publish(subCh chan interface{}, unsubCh chan interface{}, valueChs []chan interface{}) {
	for {
		fmt.Println("In for")
		select {
		case value := <-subCh:
			for i := 0; i < 10; i++ {
				value1Ch := valueChs[i]
				value1Ch <- value
			}
			fmt.Println("Sub ", value)
		case value := <-unsubCh:
			fmt.Println("Unsub ", value)
		}
	}
}

func worker(valueCh chan interface{}) {
	for value := range valueCh {
		fmt.Println("Value ", value)
	}
}
