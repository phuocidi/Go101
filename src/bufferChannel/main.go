package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var (
	running int64 = 0
)

func work() {
	atomic.AddInt64(&running, 1)
	fmt.Printf("[")
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	atomic.AddInt64(&running, -1)
	fmt.Printf("]")
}

func worker(semaphore chan bool) {
	<-semaphore
	work()
	semaphore <- true
}

func main() {
	semaphore := make(chan bool, 20)

	for i := 0; i < 1000; i++ {
		go worker(semaphore)
	}

	for i := 0; i < cap(semaphore); i++ {
		semaphore <- true
	}
	//time.Sleep(30)
}
