package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	funcoes := 5

	go manda(100, c1)
	go outra(funcoes, c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

}

func manda(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	close(c)
}

func outra(funcoes int, c1, c2 chan int) {
	var wg sync.WaitGroup

	for i := 0; i < funcoes; i++ {
		wg.Add(1)
		go func() {
			for v := range c1 {
				c2 <- trabalho(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)

}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * 1000)
	return n
}
