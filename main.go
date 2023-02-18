package main

import (
	"fmt"
	"os"

	"github.com/fabienjuif/go-adventofcode-2022/exercices"
)

func main() {
	switch os.Args[1] {
	case "1":
		exercices.RunE1()
	case "2":
		exercices.RunE2()
	default:
		panic(fmt.Sprintf("exercice not known: %v", os.Args[1]))
	}
}
