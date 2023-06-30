package main

import "fmt"

func main() {

	var p1 string

	fmt.Print("Nome:")
	_, e := fmt.Scan(&p1)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(p1)

}
