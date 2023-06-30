package main

import (
	"fmt"
	"sync"
)

func main() {

	par := make(chan int)
	impar := make(chan int)
	converg := make(chan int)

	go envia(par, impar)
	go recebe(par, impar, converg)

	for v := range converg {
		fmt.Println("valor recebido:", v)
	}
}

func envia(p, i chan int) {
	x := 100
	for n := 0; n < x; n++ {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	close(p)
	close(i)
}

func recebe(p, i, c chan int) {

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		for v := range p {
			c <- v
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for v := range i {
			c <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(c)
}
