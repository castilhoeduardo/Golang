package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}
type humanos interface {
	falar()
}

func (p *pessoa) falar() {
	p.nome = "duds"
	fmt.Println(p.nome, "falou que ele tem:", p.idade)
}
func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	p1 := pessoa{"dudu", 18}

	p1.falar()
	fmt.Println(p1.nome)
}
