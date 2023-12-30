package numbers

import "testing"

func TestAdder(t *testing.T) {
	cases := []struct {
		a, b, output int
	}{
		{1, 2, 3},
		{2, 4, 6},
		{3, 5, 8},
	}

	for _, c := range cases {
		if adder(c.a, c.b) != c.output {
			t.Errorf("Didn't work")
		}
	}
}
