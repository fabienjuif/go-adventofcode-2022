package main

import (
	"fmt"
	"os"

	"github.com/fabienjuif/go-adventofcode-2022/exercice1a"
	"github.com/fabienjuif/go-adventofcode-2022/exercice1b"
	"github.com/fabienjuif/go-adventofcode-2022/exercice2a"
	"github.com/fabienjuif/go-adventofcode-2022/exercice2b"
	"github.com/fabienjuif/go-adventofcode-2022/exercice3a"
	"github.com/fabienjuif/go-adventofcode-2022/exercice3b"
	"github.com/fabienjuif/go-adventofcode-2022/exercice4a"
	"github.com/fabienjuif/go-adventofcode-2022/exercice4b"
	"github.com/fabienjuif/go-adventofcode-2022/exercice5a"
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
	case "3a":
		exercice3a.Run()
	case "3b":
		exercice3b.Run()
	case "4a":
		exercice4a.Run()
	case "4b":
		exercice4b.Run()
	case "5a":
		exercice5a.Run()
	default:
		panic(fmt.Sprintf("exercice not known: %v", os.Args[1]))
	}
}
