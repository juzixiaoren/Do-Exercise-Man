package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func random(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	for i := 0; i < 5; i++ {
		ch <- rand.Intn(5)
	}

}
func getrandom(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		num, ok := <-ch
		if !ok {
			break
		}
		fmt.Print(num, " ")
	}

}

func main1() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)
	go random(ch, &wg)
	go getrandom(ch, &wg)
	wg.Wait()
}
