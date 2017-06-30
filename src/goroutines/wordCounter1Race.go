package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type words struct {
	found map[string]int
}

func main() {
	var wg sync.WaitGroup

	w := newWords()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}

	wg.Wait()

	fmt.Println("Words that appear more than once: ")

	for word, count := range w.found {
		if count > 1 {
			fmt.Printf(" %s: %d\n", word, count)
		}
	}
}

func newWords() *words {
	return &words{found: map[string]int{}}
}

// track the number of times seen this word
func (w *words) add(word string, n int) {

	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

// Open a file, parse its contents, and count the words that appear
func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}

	return scanner.Err()
}