package main

import (
	"fmt"
	"os"
)

func main() {

	_, e := os.Open("no-File.txt")
	if e != nil {
		fmt.Println(e)
	}
}
