package main

import (
	"fmt"
	"time"
)

func count(thing string, c chan string) {
	for i := 0; i < 10; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)

}
func main() {
	c := make(chan string)
	go count("sheep", c)
	for msg := range c {
		fmt.Println(msg)
	}
}
