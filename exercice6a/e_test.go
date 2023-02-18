package exercice6a

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
	wants := []string{"7", "5", "6", "10", "11"}

	for i, input := range inputs {
		got := process(input)
		want := wants[i]
		if want != got {
			t.Errorf(`wanted: "%v"; got "%v"`, want, got)
		}
	}
}
