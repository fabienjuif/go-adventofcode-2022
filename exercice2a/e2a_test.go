package exercice2a

import "testing"

func TestE2a(t *testing.T) {
	got := process(`A Y
B X
C Z`)

	want := `15`

	if want != got {
		t.Errorf(`wanted: "%v"; got "%v"`, want, got)
	}
}
