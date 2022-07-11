package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup
var m sync.Mutex

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

// Ex3
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

func input(ch chan string, file string) {
	defer close(ch)
	f, err := os.Open(string(file))
	if err != nil {
		fmt.Println("Can't open file")
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ch <- scanner.Text()
	}
	fmt.Println("Add to channel success")

}

func output(ch chan string) {
	for data := range ch {
		fmt.Println(data)
	}
	wg.Done()
}

func main() {
	ch := make(chan string, 40)

	//...Ex2
	// X := make(map[int]int)
	// for i := 0; i < 3; i++ {
	// 	go addToMap(X, 0, 1000)
	// }

	// time.Sleep(25 * time.Second)
	input(ch, "Data.txt")
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go output(ch)
	}
	time.Sleep(10 * time.Second)
	wg.Wait()
	fmt.Println("Done")

}
