package exercice8a

import (
	"testing"
)

func TestE(t *testing.T) {
	want := "21"
	got := process(`30373
25512
65332
33549
35390`)

	if want != got {
		t.Errorf(`wanted: "%v"; got "%v"`, want, got)
	}
}
