package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, er := json.Marshal(p1)
	if er != nil {
		log.Println("HOUVE ERRO", er)
		return
	}
	fmt.Println(string(bs))

}
