package csvlib

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Parser interface {
	Parse(val string) ([]ParsedValue, error)
	Defaults() []ParsedValue
}

type ParsedValue interface {
	I32() int32
	I64() int64
	F32() float32
	F64() float64
	String() string
	Bool() bool
	Time() time.Time
	Ifc() interface{}
	IsValid() bool
}

type RowParser struct {
	P []Parser
}

func (r *RowParser) AppendColumn(p Parser) {
	r.P = append(r.P, p)
}

type ElemNotMatch struct {
	Want, Got int
}

func (e ElemNotMatch) Error() string {
	return fmt.Sprintf("number of elements(%d) does NOT match number of parser (%d)", e.Got, e.Want)
}

var ErrNumberOfElemNotMatch = fmt.Errorf("number of elements does NOT match number of parser")

func (r *RowParser) Parse(row []string) ([]ParsedValue, error) {
	if len(row) != r.Len() {
		return nil, ElemNotMatch{r.Len(), len(row)}
	}
	var values []ParsedValue
	for i, p := range r.P {
		if row[i] == "\\N" {
			values = append(values, p.Defaults()...)
			continue
		}
		vals, err := p.Parse(row[i])
		if err != nil {
			return vals, err
		}
		values = append(values, vals...)
	}
	return values, nil
}

func (r *RowParser) Len() int {
	return len(r.P)
}

type Int32Parser struct {
	Name     string
	Optional bool  // if "" then nil
	Default  int32 // as defautl value
}

func (p Int32Parser) Parse(val string) ([]ParsedValue, error) {
	if p.Optional && val == "" {
		return p.Defaults(), nil
	}
	ret, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return nil, errors.Wrapf(err, " %s ", p.Name)
	}
	return []ParsedValue{ParsedInt32(ret)}, err

}

func (p Int32Parser) Defaults() []ParsedValue {
	return []ParsedValue{ParsedInt32(p.Default)}
}

type Int64Parser struct {
	Name string
}

func (p Int64Parser) Parse(val string) ([]ParsedValue, error) {
	ret, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, " %s ", p.Name)
	}
	return []ParsedValue{ParsedInt64(ret)}, err
}

func (Int64Parser) Defaults() []ParsedValue {
	return []ParsedValue{ParsedInt64(0)}
}

type Float32Parser struct {
	Name     string
	Optional bool
}

func (p Float32Parser) Parse(val string) ([]ParsedValue, error) {
	if val == "" && p.Optional {
		return p.Defaults(), nil
	}
	ret, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return nil, errors.Wrapf(err, " %s ", p.Name)
	}
	return []ParsedValue{ParsedFloat32(ret)}, err
}

func (Float32Parser) Defaults() []ParsedValue {
	return []ParsedValue{ParsedFloat32(0)}
}

type StringParser struct {
	Name string
}

func (StringParser) Parse(val string) ([]ParsedValue, error) {
	return []ParsedValue{ParsedString(val)}, nil
}

func (StringParser) Defaults() []ParsedValue {
	return []ParsedValue{ParsedString("")}
}

type BoolParser struct {
	Name string
}

func (p BoolParser) Parse(val string) ([]ParsedValue, error) {
	if val == "" {
		return p.Defaults(), nil
	}
	if val == "N" {
		val = "F"
	}
	ret, err := strconv.ParseBool(val)
	if err != nil {
		return p.Defaults(), errors.Wrapf(err, " %s ", p.Name)
	}
	return []ParsedValue{ParsedBool{ret, true}}, err
}

func (BoolParser) Defaults() []ParsedValue {
	return []ParsedValue{ParsedBool{false, false}}
}

type TimeParser struct {
	Name   string
	Layout string
}

func (t TimeParser) Parse(val string) ([]ParsedValue, error) {
	tm, err := time.Parse(t.Layout, val)
	if err != nil {
		return nil, errors.Wrapf(err, " %s ", t.Name)
	}
	return []ParsedValue{ParsedTime(tm)}, err
}

func (TimeParser) Defaults() []ParsedValue {
	return []ParsedValue{ParsedTime(time.Unix(0, 0))}
}
