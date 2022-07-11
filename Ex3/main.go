package main

import (
	"log"
	"sync"
)

func errFunc() {
	var mutex sync.Mutex
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j < 10000; j++ {
				mutex.Lock()
				if _, ok := m[j]; ok {
					delete(m, j)
					continue
				}
				m[j] = j * 10
				mutex.Unlock()
			}
		}()
	}

	log.Print("done")
}

func main() {
	errFunc()
}
