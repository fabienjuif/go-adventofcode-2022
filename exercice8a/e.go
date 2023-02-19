package exercice8a

// ----
// This exercice could have been done with simple loops
// But I wanted a pretext to use go routines + chanels + wait group!
// ----

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func Run() {
	fmt.Println("exercice 8a")
	fmt.Println(process(input))
}

func process(text string) string {
	lines := strings.Split(text, "\n")

	w := len(lines[0])
	h := len(lines)
	forest := NewForest(w, h)

	for x := 0; x < w; x += 1 {
		for y := 0; y < h; y += 1 {
			forest.GetTree(x, y).Tall = StringToInt(lines[y][x : x+1])
		}
	}

	forest.Roam()

	// fmt.Println(forest)

	return fmt.Sprintf("%d", forest.CountVisibleTrees())
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

type Forest struct {
	Matrice []*Tree
	Width   int
	Height  int
}

func (f Forest) CountVisibleTrees() uint {
	c := uint(0)
	for _, t := range f.Matrice {
		if t.Visibility {
			c += 1
		}
	}
	return c
}

func (f Forest) String() string {
	str := ""
	for y := 0; y < f.Height; y += 1 {
		for x := 0; x < f.Width; x += 1 {
			if x == 0 {
				str += "|"
			}
			tree := f.GetTree(x, y)
			visible := " "
			if tree.Visibility {
				visible = "*"
			}
			str += fmt.Sprintf("%s%d|", visible, tree.Tall)
		}
		str += "\n"
	}
	return str
}

func NewForest(w, h int) *Forest {
	matrice := make([]*Tree, h*w)

	for i := 0; i < w*h; i += 1 {
		x := i % w
		y := i / w
		matrice[i] = NewTree(x, y, 0)
	}

	return &Forest{
		Matrice: matrice,
		Width:   w,
		Height:  h,
	}
}

func (f *Forest) GetTree(x int, y int) *Tree {
	return f.Matrice[x+y*f.Width]
}

func (f *Forest) Roam() {
	var wg sync.WaitGroup
	ch := make(chan RoamEvent, 10)
	chCapturing := make(chan bool, 1)

	// roaming
	wg.Add(4)
	go f.RoamToOuest(ch, &wg)
	go f.RoamToEast(ch, &wg)
	go f.RoamToNorth(ch, &wg)
	go f.RoamToSouth(ch, &wg)

	// capturing roaming
	go func() {
		for e := range ch {
			e.Tree.Visibility = e.Tree.Visibility || e.Visibility
		}
		close(chCapturing)
	}()

	wg.Wait()
	close(ch)
	<-chCapturing
}

func (f *Forest) RoamToOuest(ch chan RoamEvent, wg *sync.WaitGroup) {
	defer wg.Done()
	start := f.Width - 1
	end := 0

	for y := 0; y < f.Height; y += 1 {
		maxTall := -1
		for x := start; x >= end; x -= 1 {
			maxTall = f.processTreeVisibility(ch, maxTall, x, y)
		}
	}
}

func (f *Forest) RoamToEast(ch chan RoamEvent, wg *sync.WaitGroup) {
	defer wg.Done()
	start := 0
	end := f.Width - 1

	for y := 0; y < f.Height; y += 1 {
		maxTall := -1
		for x := start; x < end; x += 1 {
			maxTall = f.processTreeVisibility(ch, maxTall, x, y)
		}
	}
}

func (f *Forest) RoamToNorth(ch chan RoamEvent, wg *sync.WaitGroup) {
	defer wg.Done()
	start := f.Height - 1
	end := 0

	for x := 0; x < f.Width; x += 1 {
		maxTall := -1
		for y := start; y >= end; y -= 1 {
			maxTall = f.processTreeVisibility(ch, maxTall, x, y)
		}
	}
}

func (f *Forest) RoamToSouth(ch chan RoamEvent, wg *sync.WaitGroup) {
	defer wg.Done()
	start := 0
	end := f.Height - 1

	for x := 0; x < f.Width; x += 1 {
		maxTall := -1
		for y := start; y < end; y += 1 {
			maxTall = f.processTreeVisibility(ch, maxTall, x, y)
		}
	}
}

func (f *Forest) processTreeVisibility(ch chan RoamEvent, maxTall int, x int, y int) int {
	curTree := f.GetTree(x, y)
	curMaxTall := maxTall

	visibility := false
	if curTree.Tall > curMaxTall {
		visibility = true
		curMaxTall = curTree.Tall
	}

	ch <- RoamEvent{
		Tree:       curTree,
		Visibility: visibility,
	}

	return curMaxTall
}

type RoamEvent struct {
	Visibility bool
	Tree       *Tree
}

type RoamEventType struct {
	slug string
}

func (r RoamEventType) String() string {
	return r.slug
}

var (
	RoamEventOuest = RoamEventType{"ouest"}
	RoamEventEast  = RoamEventType{"est"}
	RoamEventNorth = RoamEventType{"north"}
	RoamEventSouth = RoamEventType{"south"}
)

type Tree struct {
	X          int
	Y          int
	Tall       int
	Visibility bool
}

func NewTree(X int, Y int, Tall int) *Tree {
	t := &Tree{
		X:          X,
		Y:          Y,
		Tall:       Tall,
		Visibility: false,
	}

	return t
}
