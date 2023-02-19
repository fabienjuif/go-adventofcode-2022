package exercice8b

import (
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("exercice 8b")
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

	// fmt.Println(forest)

	return fmt.Sprintf("%d", forest.FindHightScenicScore())
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

func (f Forest) String() string {
	str := ""
	for y := 0; y < f.Height; y += 1 {
		for x := 0; x < f.Width; x += 1 {
			if x == 0 {
				str += "|"
			}
			tree := f.GetTree(x, y)
			str += fmt.Sprintf("%d|", tree.Tall)
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

func (f *Forest) FindHightScenicScore() int {
	highest := 0

	for _, t := range f.Matrice {
		scenicR := 0
		scenicL := 0
		scenicT := 0
		scenicB := 0

		for x := t.X + 1; x < f.Width; x += 1 {
			n := f.GetTree(x, t.Y)
			if n == nil {
				break
			}
			scenicR += 1
			if n.Tall >= t.Tall {
				break
			}
		}

		for y := t.Y + 1; y < f.Height; y += 1 {
			n := f.GetTree(t.X, y)
			if n == nil {
				break
			}
			scenicB += 1
			if n.Tall >= t.Tall {
				break
			}
		}

		for x := t.X - 1; x >= 0; x -= 1 {
			n := f.GetTree(x, t.Y)
			if n == nil {
				break
			}
			scenicL += 1
			if n.Tall >= t.Tall {
				break
			}
		}

		for y := t.Y - 1; y >= 0; y -= 1 {
			n := f.GetTree(t.X, y)
			if n == nil {
				break
			}
			scenicT += 1
			if n.Tall >= t.Tall {
				break
			}
		}

		scenic := scenicR * scenicL * scenicT * scenicB
		if scenic > highest {
			highest = scenic
		}
	}

	return highest
}

func (f *Forest) GetTree(x int, y int) *Tree {
	if x < 0 || x >= f.Width || y < 0 || y >= f.Height {
		return nil
	}
	idx := x + y*f.Width
	return f.Matrice[idx]
}

type Tree struct {
	X    int
	Y    int
	Tall int
}

func NewTree(X int, Y int, Tall int) *Tree {
	t := &Tree{
		X:    X,
		Y:    Y,
		Tall: Tall,
	}

	return t
}
