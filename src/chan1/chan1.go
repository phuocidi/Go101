package main

import (
	"fmt"
	"time"
)

func emit(chanelChan chan chan string, done chan bool) {
	wordsChan := make(chan string)
	chanelChan <- wordsChan
	defer close(wordsChan)

	words := []string{"Tran", "Huu", "Phuoc", "is", "so", "dzach"}

	i := 0
	t := time.NewTimer(3 * time.Second)

	for {
		select {
		case wordsChan <- words[i]:
			i += 1
			if i == len(words) {
				i = 0
			}

		case <-done:
			fmt.Printf("[%v] CLose", done)
			done <- true
			return

		case <-t.C:
			return
		}
	}
}

func main() {
	chanelChan := make(chan chan string)
	doneCh := make(chan bool)

	go emit(chanelChan, doneCh)

	wordCh := <-chanelChan

	for word := range wordCh {
		fmt.Printf("%s\n", word)
	}

}
