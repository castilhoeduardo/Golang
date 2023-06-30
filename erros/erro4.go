package main

import (
	"log"
	"os"
)

func main() {

	_, e := os.Open("no-File.txt")
	if e != nil {
		log.Println(e)
	}
}
