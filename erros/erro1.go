package main

import "fmt"

func main() {

	var p1, p2, p3 string
	fmt.Println("Name: ")
	_, e := fmt.Scan(&p1)
	if e != nil {
		panic(e)
	}
	fmt.Print("Fav food: ")
	_, e = fmt.Scan(&p2)
	if e != nil {
		panic(e)
	}
	fmt.Print("Fav sport: ")
	_, e = fmt.Scan(&p3)
	if e != nil {
		panic(e)
	}
	fmt.Println(p1, p2, p3)
}
