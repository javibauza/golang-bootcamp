package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)
	done := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ticker.C:
				ch <- "ping"
			case <-done:
				wg.Done()
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Printf("Message: %s\n", msg)
			case <-done:
				wg.Done()
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)
	done <- true
	done <- true
	wg.Wait()
}
