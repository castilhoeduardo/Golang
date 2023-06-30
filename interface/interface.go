package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) falar() {
	fmt.Println(p.nome, "diz: oi, ele tem:", p.idade, "anos.")
}

type humanos interface {
	falar()
}

func dizeralgo(h humanos) {
	h.falar()
}

func main() {

	p1 := pessoa{"Dudu", 18}
	p1.falar()

	dizeralgo(&p1)
}
