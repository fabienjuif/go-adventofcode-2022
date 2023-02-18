package exercice3a

import (
	"testing"
)

func TestE3a(t *testing.T) {
	got := process(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)

	want := `157`

	if want != got {
		t.Errorf(`wanted: "%v"; got "%v"`, want, got)
	}
}

func TestByteToScore(t *testing.T) {
	inputs := []byte{'a', 'z', 'A', 'Z'}
	wants := []int{1, 26, 27, 52}

	for idx, input := range inputs {
		want := wants[idx]
		got := ByteToScore(input)

		if want != got {
			t.Errorf(`wanted: "%v"; got "%v"`, want, got)
		}
	}
}
