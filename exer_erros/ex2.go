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

	bs, e := toJSON(p1)
	if e != nil {
		log.Fatalln("caiu no saco")
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, e := json.Marshal(a)
	if e != nil {
		return []byte{}, fmt.Errorf("Ei deu erro nego")
	}
	return bs, e
}
