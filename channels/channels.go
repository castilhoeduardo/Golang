package main

import (
	"fmt"
)

func main() {

	c1 := make(chan int)

	go func() {
		c1 <- 42
		close(c1)
	}()

	v, ok := <-c1

	fmt.Println(v, ok)
}
