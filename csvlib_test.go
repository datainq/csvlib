package csvlib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

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

func TestRowParser(t *testing.T) {
	parser := RowParser{P: []Parser{
		StringParser{},
		Int64Parser{},
		SkipParser{},
	},
	}
	row, err := parser.Parse([]string{"123", "321", "333"})
	assert.NoError(t, err)
	assert.Len(t, row, 2)
	assert.Equal(t,
		[]interface{}{"123", int64(321)},
		[]interface{}{row[0].String(), row[1].I64()})
}