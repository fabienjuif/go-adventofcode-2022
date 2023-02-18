package exercice5a

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("exercice 5a")
	fmt.Println(process(input))
}

func process(text string) string {
	commandRegexp := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	lines := strings.Split(text, "\n")

	stackCount := 0
	var stacks []*Stack = []*Stack{}
	init := true
	for idx, line := range lines {
		if line == "" {
			continue
		}
		if init {
			if idx == 0 {
				stackCount = GetStackCount(line)
				stacks = make([]*Stack, stackCount)
				for sIdx := 0; sIdx < stackCount; sIdx += 1 {
					stacks[sIdx] = &Stack{Crates: make([]byte, 0, 10)}
				}
			}

			if line[1] == '1' {
				init = false
				continue
			}

			for iS, s := range stacks {
				crate := line[1+iS*4]
				if crate != ' ' {
					s.Unshift(crate)
				}
			}
		} else {
			matches := commandRegexp.FindStringSubmatch(line)
			howMany := StringToInt(matches[1])
			from := StringToInt(matches[2])
			to := StringToInt(matches[3])

			stacks[from-1].MoveInto(howMany, stacks[to-1])
		}
	}

	head := make([]byte, stackCount)
	for i, stack := range stacks {
		head[i] = stack.Crates[len(stack.Crates)-1]
	}

	return string(head)
}

func PrintStacks(stacks []*Stack) {
	for i, s := range stacks {
		fmt.Printf("Stack %d: ", i+1)
		l := len(s.Crates)
		for j := 0; j < l; j += 1 {
			fmt.Printf("%s", string([]byte{s.Crates[j]}))
		}
		fmt.Printf("\n")
	}
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func GetStackCount(line string) int {
	return (len(line)-7)/4 + 2
}

type Stack struct {
	Crates []byte
}

func (s *Stack) Push(crate byte) {
	s.Crates = append(s.Crates, crate)
}

func (s *Stack) Unshift(crate byte) {
	l := len(s.Crates)
	newCrates := make([]byte, l+1)
	newCrates[0] = crate
	for i := 1; i < l+1; i += 1 {
		newCrates[i] = s.Crates[i-1]
	}
	s.Crates = newCrates
}

func (s *Stack) Pop() byte {
	index := len(s.Crates) - 1
	r := s.Crates[index]
	s.Crates = s.Crates[:index]
	return r
}

func (s *Stack) MoveInto(size int, o *Stack) {
	for i := 0; i < size; i += 1 {
		o.Push(s.Pop())
	}
}

var input = `[N]     [Q]         [N]            
[R]     [F] [Q]     [G] [M]        
[J]     [Z] [T]     [R] [H] [J]    
[T] [H] [G] [R]     [B] [N] [T]    
[Z] [J] [J] [G] [F] [Z] [S] [M]    
[B] [N] [N] [N] [Q] [W] [L] [Q] [S]
[D] [S] [R] [V] [T] [C] [C] [N] [G]
[F] [R] [C] [F] [L] [Q] [F] [D] [P]
 1   2   3   4   5   6   7   8   9 

move 3 from 9 to 4
move 2 from 5 to 2
move 8 from 1 to 9
move 4 from 7 to 1
move 5 from 3 to 8
move 3 from 3 to 7
move 11 from 8 to 3
move 7 from 3 to 6
move 2 from 5 to 9
move 3 from 1 to 6
move 6 from 2 to 4
move 6 from 7 to 5
move 1 from 6 to 1
move 1 from 9 to 4
move 16 from 4 to 9
move 2 from 1 to 2
move 1 from 4 to 6
move 1 from 3 to 7
move 2 from 2 to 4
move 1 from 7 to 9
move 22 from 9 to 8
move 1 from 6 to 3
move 18 from 8 to 5
move 3 from 8 to 2
move 3 from 2 to 9
move 13 from 6 to 7
move 1 from 6 to 7
move 4 from 3 to 6
move 2 from 6 to 3
move 2 from 3 to 8
move 3 from 7 to 8
move 14 from 5 to 2
move 3 from 2 to 5
move 2 from 8 to 4
move 4 from 8 to 6
move 4 from 6 to 4
move 11 from 5 to 2
move 3 from 9 to 2
move 7 from 2 to 3
move 11 from 7 to 2
move 1 from 5 to 7
move 5 from 6 to 8
move 30 from 2 to 7
move 23 from 7 to 2
move 4 from 3 to 4
move 3 from 9 to 6
move 4 from 8 to 2
move 1 from 8 to 2
move 2 from 7 to 9
move 4 from 2 to 3
move 1 from 5 to 9
move 6 from 4 to 7
move 5 from 3 to 6
move 1 from 3 to 6
move 1 from 9 to 2
move 16 from 2 to 5
move 7 from 7 to 6
move 9 from 2 to 1
move 2 from 1 to 4
move 8 from 5 to 3
move 5 from 7 to 4
move 1 from 9 to 8
move 9 from 3 to 6
move 25 from 6 to 8
move 2 from 9 to 5
move 3 from 4 to 2
move 7 from 4 to 1
move 1 from 8 to 7
move 6 from 5 to 2
move 11 from 8 to 5
move 1 from 7 to 9
move 10 from 1 to 2
move 6 from 5 to 1
move 1 from 4 to 2
move 13 from 8 to 1
move 17 from 1 to 2
move 5 from 1 to 9
move 1 from 8 to 4
move 1 from 1 to 3
move 1 from 3 to 6
move 1 from 9 to 3
move 2 from 4 to 5
move 1 from 4 to 8
move 1 from 9 to 1
move 8 from 5 to 7
move 1 from 8 to 1
move 7 from 7 to 6
move 2 from 1 to 2
move 1 from 3 to 6
move 2 from 5 to 4
move 8 from 2 to 1
move 1 from 9 to 7
move 1 from 5 to 1
move 2 from 7 to 3
move 2 from 3 to 7
move 2 from 7 to 8
move 2 from 1 to 5
move 3 from 9 to 2
move 2 from 8 to 9
move 1 from 9 to 2
move 1 from 9 to 8
move 1 from 8 to 7
move 6 from 6 to 5
move 1 from 6 to 2
move 2 from 4 to 5
move 2 from 6 to 8
move 1 from 7 to 1
move 2 from 8 to 4
move 11 from 2 to 5
move 18 from 5 to 6
move 6 from 2 to 6
move 10 from 2 to 7
move 1 from 4 to 3
move 7 from 2 to 8
move 7 from 1 to 4
move 6 from 7 to 8
move 2 from 7 to 3
move 8 from 4 to 7
move 1 from 1 to 3
move 1 from 2 to 1
move 4 from 7 to 1
move 4 from 1 to 3
move 2 from 3 to 9
move 2 from 5 to 4
move 1 from 2 to 1
move 2 from 1 to 5
move 1 from 3 to 1
move 2 from 5 to 2
move 1 from 2 to 6
move 5 from 7 to 4
move 1 from 1 to 2
move 10 from 8 to 1
move 2 from 2 to 7
move 2 from 7 to 1
move 1 from 7 to 9
move 1 from 5 to 7
move 3 from 8 to 7
move 3 from 3 to 6
move 3 from 7 to 1
move 5 from 1 to 4
move 1 from 7 to 6
move 22 from 6 to 3
move 2 from 6 to 2
move 19 from 3 to 4
move 15 from 4 to 8
move 9 from 8 to 4
move 5 from 6 to 8
move 2 from 2 to 8
move 2 from 9 to 4
move 7 from 1 to 5
move 1 from 1 to 3
move 1 from 9 to 7
move 5 from 8 to 3
move 4 from 8 to 1
move 5 from 1 to 5
move 10 from 4 to 3
move 3 from 4 to 2
move 2 from 8 to 3
move 12 from 4 to 8
move 1 from 7 to 6
move 3 from 2 to 9
move 2 from 4 to 5
move 5 from 3 to 7
move 1 from 7 to 2
move 1 from 1 to 6
move 1 from 7 to 2
move 15 from 3 to 8
move 10 from 5 to 6
move 3 from 7 to 8
move 1 from 5 to 8
move 1 from 2 to 3
move 7 from 6 to 1
move 3 from 5 to 3
move 5 from 3 to 5
move 3 from 5 to 4
move 2 from 4 to 9
move 2 from 3 to 5
move 14 from 8 to 5
move 1 from 9 to 1
move 16 from 5 to 3
move 16 from 3 to 6
move 2 from 9 to 8
move 21 from 6 to 7
move 2 from 1 to 7
move 1 from 2 to 7
move 4 from 1 to 7
move 1 from 4 to 7
move 16 from 8 to 5
move 20 from 7 to 4
move 1 from 9 to 8
move 1 from 7 to 4
move 3 from 8 to 6
move 1 from 9 to 1
move 2 from 1 to 4
move 2 from 5 to 2
move 5 from 4 to 7
move 1 from 6 to 9
move 11 from 7 to 6
move 2 from 7 to 5
move 12 from 6 to 2
move 13 from 2 to 1
move 1 from 2 to 3
move 1 from 8 to 4
move 6 from 4 to 1
move 1 from 6 to 7
move 7 from 4 to 9
move 8 from 9 to 3
move 2 from 8 to 3
move 10 from 5 to 4
move 11 from 1 to 8
move 1 from 1 to 3
move 5 from 1 to 8
move 8 from 5 to 6
move 13 from 8 to 9
move 12 from 3 to 5
move 12 from 5 to 9
move 1 from 7 to 9
move 1 from 1 to 2
move 1 from 1 to 4
move 3 from 8 to 5
move 1 from 2 to 5
move 1 from 4 to 8
move 5 from 6 to 3
move 1 from 8 to 4
move 13 from 4 to 7
move 3 from 7 to 6
move 1 from 1 to 4
move 4 from 4 to 2
move 1 from 6 to 3
move 2 from 5 to 9
move 2 from 5 to 9
move 1 from 4 to 8
move 6 from 9 to 4
move 22 from 9 to 2
move 8 from 7 to 4
move 7 from 2 to 1
move 3 from 3 to 8
move 2 from 6 to 7
move 14 from 4 to 2
move 2 from 6 to 1
move 1 from 8 to 7
move 3 from 3 to 9
move 1 from 8 to 4
move 3 from 1 to 9
move 3 from 9 to 3
move 31 from 2 to 8
move 8 from 8 to 4
move 1 from 9 to 1
move 9 from 4 to 5
move 7 from 5 to 6
move 2 from 5 to 1
move 1 from 2 to 1
move 1 from 7 to 9
move 1 from 2 to 9
move 2 from 6 to 4
move 4 from 7 to 4
move 4 from 9 to 8
move 6 from 4 to 1
move 1 from 3 to 2
move 1 from 3 to 6
move 1 from 9 to 2
move 2 from 2 to 4
move 1 from 9 to 1
move 1 from 3 to 1
move 17 from 1 to 9
move 4 from 6 to 2
move 1 from 9 to 7
move 4 from 9 to 7
move 1 from 8 to 4
move 3 from 7 to 6
move 1 from 4 to 9
move 10 from 8 to 5
move 6 from 6 to 5
move 1 from 7 to 2
move 1 from 1 to 4
move 1 from 4 to 5
move 9 from 8 to 3
move 4 from 3 to 9
move 1 from 4 to 6
move 1 from 6 to 5
move 1 from 4 to 8
move 2 from 3 to 8
move 1 from 3 to 8
move 3 from 8 to 9
move 5 from 2 to 9
move 1 from 7 to 9
move 5 from 8 to 7
move 3 from 8 to 4
move 2 from 8 to 5
move 24 from 9 to 7
move 1 from 3 to 5
move 2 from 9 to 4
move 22 from 7 to 9
move 1 from 3 to 4
move 6 from 4 to 6
move 4 from 6 to 7
move 10 from 5 to 3
move 8 from 3 to 5
move 2 from 3 to 4
move 2 from 4 to 6
move 10 from 7 to 3
move 21 from 9 to 1
move 2 from 3 to 4
move 4 from 3 to 8
move 2 from 4 to 8
move 1 from 7 to 8
move 4 from 6 to 8
move 3 from 5 to 4
move 8 from 8 to 2
move 18 from 1 to 6
move 3 from 4 to 1
move 1 from 2 to 8
move 5 from 1 to 8
move 3 from 3 to 6
move 3 from 2 to 9
move 3 from 8 to 1
move 11 from 5 to 2
move 3 from 8 to 7
move 10 from 2 to 9
move 1 from 7 to 9
move 3 from 8 to 1
move 2 from 7 to 8
move 6 from 9 to 5
move 4 from 2 to 8
move 8 from 5 to 8
move 1 from 3 to 7
move 2 from 5 to 6
move 3 from 1 to 6
move 2 from 1 to 6
move 4 from 9 to 8
move 4 from 9 to 8
move 1 from 9 to 4
move 9 from 6 to 9
move 16 from 6 to 9
move 1 from 4 to 7
move 1 from 2 to 9
move 5 from 8 to 5
move 4 from 5 to 1
move 6 from 1 to 7
move 12 from 8 to 4
move 5 from 8 to 1
move 6 from 9 to 3
move 1 from 1 to 6
move 2 from 5 to 8
move 12 from 4 to 7
move 2 from 8 to 4
move 1 from 4 to 8
move 1 from 7 to 6
move 1 from 4 to 6
move 14 from 7 to 1
move 3 from 3 to 2
move 7 from 9 to 7
move 3 from 3 to 5
move 15 from 1 to 2
move 2 from 5 to 9
move 1 from 8 to 9
move 16 from 9 to 1
move 1 from 5 to 9
move 5 from 6 to 2
move 12 from 7 to 2
move 20 from 2 to 6
move 10 from 2 to 6
move 11 from 1 to 7
move 2 from 7 to 4
move 2 from 2 to 5
move 1 from 2 to 3
move 2 from 5 to 6
move 1 from 9 to 5
move 1 from 5 to 9
move 25 from 6 to 7
move 25 from 7 to 6
move 1 from 3 to 1
move 1 from 2 to 5
move 1 from 4 to 3
move 33 from 6 to 3
move 1 from 9 to 5
move 2 from 3 to 5
move 28 from 3 to 9
move 5 from 1 to 9
move 4 from 1 to 8
move 2 from 3 to 2
move 2 from 8 to 1
move 1 from 4 to 6
move 3 from 5 to 3
move 1 from 2 to 4
move 2 from 2 to 8
move 1 from 6 to 5
move 30 from 9 to 2
move 7 from 2 to 6
move 1 from 1 to 3
move 1 from 1 to 7
move 1 from 5 to 6
move 1 from 5 to 4
move 5 from 7 to 4
move 4 from 7 to 3
move 1 from 3 to 7
move 3 from 8 to 7
move 8 from 3 to 1
move 3 from 1 to 7
move 4 from 1 to 4
move 3 from 9 to 8
move 8 from 6 to 2
move 18 from 2 to 6
move 6 from 7 to 2
move 1 from 1 to 7
move 3 from 4 to 7
move 5 from 4 to 8
move 2 from 8 to 7
move 7 from 2 to 5
move 5 from 2 to 7
move 10 from 7 to 9
move 5 from 5 to 9
move 1 from 3 to 9
move 5 from 2 to 6
move 3 from 7 to 9
move 2 from 5 to 6
move 2 from 2 to 9
move 2 from 8 to 7
move 1 from 4 to 5
move 8 from 9 to 2
move 5 from 6 to 7
move 4 from 9 to 1
move 4 from 8 to 9
move 12 from 9 to 1
move 16 from 1 to 4
move 12 from 6 to 2
move 3 from 7 to 6
move 3 from 7 to 3
move 1 from 9 to 4
move 12 from 4 to 5
move 2 from 4 to 3
move 1 from 7 to 1
move 4 from 4 to 3
move 1 from 8 to 2
move 6 from 3 to 1
move 1 from 1 to 8
move 7 from 2 to 5
move 1 from 8 to 1
move 4 from 5 to 4
move 5 from 5 to 9
move 1 from 3 to 8
move 1 from 9 to 7
move 1 from 8 to 1
move 4 from 5 to 6
move 5 from 5 to 9
move 7 from 9 to 5
move 11 from 6 to 4
move 1 from 9 to 4
move 1 from 9 to 1
move 1 from 7 to 2
move 9 from 4 to 3
move 5 from 1 to 9
move 3 from 5 to 1
move 5 from 9 to 8
move 8 from 3 to 1
move 2 from 5 to 3
move 7 from 2 to 5
move 1 from 6 to 4
move 3 from 5 to 9
move 3 from 6 to 9
move 3 from 2 to 9
move 5 from 3 to 6
move 1 from 9 to 5
move 4 from 8 to 3
move 1 from 8 to 4
move 8 from 1 to 3
move 12 from 3 to 7
move 1 from 2 to 4
move 3 from 2 to 8
move 6 from 7 to 6
move 4 from 5 to 7
move 5 from 9 to 7
move 2 from 9 to 2
move 1 from 9 to 5
move 4 from 5 to 1
move 1 from 5 to 4
move 14 from 7 to 6
move 1 from 1 to 7
move 10 from 4 to 5
move 4 from 1 to 2
move 1 from 4 to 6
move 1 from 7 to 4
move 17 from 6 to 8
move 1 from 5 to 7
move 10 from 5 to 4
move 1 from 2 to 6
move 4 from 2 to 6
move 13 from 6 to 1
move 9 from 4 to 3
move 2 from 2 to 4
move 1 from 6 to 7
move 1 from 4 to 3
move 8 from 3 to 5
move 1 from 3 to 4
move 17 from 1 to 3
move 15 from 3 to 7
move 3 from 4 to 1
move 6 from 8 to 9
move 6 from 9 to 1
move 2 from 3 to 1
move 2 from 5 to 2
move 6 from 7 to 6
move 3 from 6 to 9`
