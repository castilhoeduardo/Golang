package main

import "fmt"

func bubblesort(x []int) []int {
	for i := range x {
		for j := 0; j < len(x)-i-1; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
	return x
}
func main() {
	x := []int{3, 7, 8, 6, 9, 2}
	bubblesort(x)
	fmt.Println(x)

}
