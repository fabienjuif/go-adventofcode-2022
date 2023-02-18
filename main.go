package main

import (
	"fmt"
	"os"

	"github.com/fabienjuif/go-adventofcode-2022/exercice1a"
	"github.com/fabienjuif/go-adventofcode-2022/exercice1b"
	"github.com/fabienjuif/go-adventofcode-2022/exercice2a"
	"github.com/fabienjuif/go-adventofcode-2022/exercice2b"
)

func main() {
	switch os.Args[1] {
	case "1a":
		exercice1a.Run()
	case "1b":
		exercice1b.Run()
	case "2a":
		exercice2a.Run()
	case "2b":
		exercice2b.Run()
	default:
		panic(fmt.Sprintf("exercice not known: %v", os.Args[1]))
	}
}
