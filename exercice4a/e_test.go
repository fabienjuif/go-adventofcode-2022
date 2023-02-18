package exercice4a

import (
	"testing"
)

func TestE4a(t *testing.T) {
	got := process(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)

	want := `2`

	if want != got {
		t.Errorf(`wanted: "%v"; got "%v"`, want, got)
	}
}
