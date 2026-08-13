package main

import (
	"sync"
	"testing"
)

func Test_random(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)
	go random(ch, &wg)
	go getrandom(ch, &wg)
	wg.Wait()
	t.Logf("goroutine done")
}
