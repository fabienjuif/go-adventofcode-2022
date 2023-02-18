package main

import (
	"fmt"
	"os"

	"github.com/fabienjuif/go-adventofcode-2022/exercices"
)

func main() {
	switch os.Args[1] {
	case "1a":
		exercices.RunE1a()
	case "1b":
		exercices.RunE1b()
	case "2":
		exercices.RunE2()
	default:
		panic(fmt.Sprintf("exercice not known: %v", os.Args[1]))
	}
}
