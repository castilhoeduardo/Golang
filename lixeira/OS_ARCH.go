package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Meu sistema operacional é:", runtime.GOOS, "/ E a arquitetura da minha cpu é:", runtime.GOARCH)
}
