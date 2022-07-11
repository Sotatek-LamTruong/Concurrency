package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex

func addToMap(X map[int]int, from int, to int) {
	var count int = 0
	for i := from; i < to; i++ {
		m.Lock()
		X[i] = i
		count++
		m.Unlock()
		// fmt.Println(fmt.Sprintf("Value of key %d is %d", i, X[i]))
	}
	fmt.Println(count)
}

func main() {
	X := make(map[int]int)
	go addToMap(X, 0, 1000)
	go addToMap(X, 1000, 2000)
	go addToMap(X, 2000, 3000)
	time.Sleep(25 * time.Second)
}
