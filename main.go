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
	"github.com/fabienjuif/go-adventofcode-2022/exercice5b"
	"github.com/fabienjuif/go-adventofcode-2022/exercice6a"
	"github.com/fabienjuif/go-adventofcode-2022/exercice6b"
	"github.com/fabienjuif/go-adventofcode-2022/exercice7a"
	"github.com/fabienjuif/go-adventofcode-2022/exercice7b"
	"github.com/fabienjuif/go-adventofcode-2022/exercice8a"
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
	case "5b":
		exercice5b.Run()
	case "6a":
		exercice6a.Run()
	case "6b":
		exercice6b.Run()
	case "7a":
		exercice7a.Run()
	case "7b":
		exercice7b.Run()
	case "8a":
		exercice8a.Run()
	default:
		panic(fmt.Sprintf("exercice not known: %v", os.Args[1]))
	}
}
