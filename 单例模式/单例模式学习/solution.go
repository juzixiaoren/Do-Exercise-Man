package main

import (
	"fmt"
	"sync"
)

type TClass struct {
	name string
}

var t *TClass

func initT() {
	t = &TClass{
		name: "AAA",
	}
}

var once sync.Once

func singleMode() *TClass {
	once.Do(initT)
	return t
}

func main() {
	t1 := singleMode()
	t2 := singleMode()
	if t1 == t2 {
		fmt.Println("单例模式成功")
	}
}
