package exercice5a

import (
	"testing"
)

func TestE(t *testing.T) {
	got := process(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)

	want := `CMZ`

	if want != got {
		t.Errorf(`wanted: "%v"; got "%v"`, want, got)
	}
}
