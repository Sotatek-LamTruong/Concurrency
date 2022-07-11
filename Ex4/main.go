package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

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

	input(ch, "Data.txt")
	// var index int = 0
	// wg.Add(3)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go output(ch)
	}
	time.Sleep(10 * time.Second)
	wg.Wait()
	// close(ch)
	fmt.Println("Done")

}
