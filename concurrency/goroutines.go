package main

import (
	"fmt"
	"sync"
	"time"
)

func saywithtime(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		saywithtime("Hello")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		saywithtime("World")
		wg.Done()
	}()
	wg.Wait()
}
