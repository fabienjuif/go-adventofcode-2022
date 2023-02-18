package exercice6b

import (
	"testing"
)

func TestE(t *testing.T) {
	inputs := []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	}
	wants := []string{"19", "23", "23", "29", "26"}

	for i, input := range inputs {
		got := process(input)
		want := wants[i]
		if want != got {
			t.Errorf(`wanted: "%v"; got "%v"`, want, got)
		}
	}
}
