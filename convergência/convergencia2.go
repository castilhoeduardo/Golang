package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c1 := converge(trabalho("Pri"), trabalho("Seg"))

	for x := 0; x < 16; x++ {
		fmt.Println(<-c1)
	}
}

func trabalho(s string) chan string {
	c1 := make(chan string)

	go func(s string, c chan string) {
		for i := 1; ; i++ {
			c <- fmt.Sprintf("Função %v diz: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(s, c1)
	return c1
}
func converge(x, y chan string) chan string {
	novo := make(chan string)
	go func() {
		for {
			novo <- <-x
		}
	}()
	go func() {
		for {
			novo <- <-y
		}
	}()
	return novo
}
