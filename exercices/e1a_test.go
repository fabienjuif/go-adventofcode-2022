package exercices

import "testing"

func TestE1a(t *testing.T) {
	got := processE1a(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)

	want := `24000`

	if want != got {
		t.Errorf(`wanted: "%v"; got "%v"`, want, got)
	}
}
