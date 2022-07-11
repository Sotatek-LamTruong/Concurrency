package main

import (
	"log"
	"sync"
	"time"
)

func chanRoutine() {
	var wg sync.WaitGroup
	wg.Add(1)
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		wg.Done()
	}()
	wg.Wait()
	log.Print("hello 2")
}

func main() {
	chanRoutine()
}
