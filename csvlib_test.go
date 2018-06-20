package csvlib

import "testing"

func TestBoolParser(t *testing.T) {
	bp := BoolParser{}

	for _, s := range []struct {
		s    string
		want bool
	}{
		{"true", true},
		{"1", true},
		{"false", false},
		{"0", false},
	} {
		v, err := bp.Parse(s.s)
		if err != nil {
			t.Errorf("got err: %s", err)
			t.FailNow()
		}
		if len(v) != 1 {
			t.Errorf("got 1 value, got: %d", len(v))
			t.FailNow()
		}
		if w := v[0].Bool(); s.want != w {
			t.Errorf("want: %t, got: %t", s.want, w)
		}
	}
}
