package main

import (
	"fmt"
	"time"
)

func numeros(done chan<- bool) {
	for i := 1; i <= 26; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}
	done <- true
}

func letras(done chan<- bool) {
	for l := 'a'; l <= 'z'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 230)
	}
	done <- true
}

func main() {
	c1 := make(chan bool)
	c2 := make(chan bool)

	go numeros(c1)
	go letras(c2)

	<-c1
	<-c2

	fmt.Println("\nFIM DA EXECUÇÃO!")
}
