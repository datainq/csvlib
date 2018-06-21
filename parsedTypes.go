package csvlib

import (
	"time"
	"github.com/ericlagergren/decimal"
)

type ParsedDefault struct{}

func (ParsedDefault) I32() int32 {
	panic("method not available")
}

func (ParsedDefault) I64() int64 {
	panic("method not available")
}

func (ParsedDefault) F32() float32 {
	panic("method not available")
}

func (ParsedDefault) F64() float64 {
	panic("method not available")
}

func (ParsedDefault) String() string {
	panic("method not available")
}

func (p ParsedDefault) Bool() bool {
	panic("method not available")
}

func (ParsedDefault) Time() time.Time {
	panic("method not available")
}

func (ParsedDefault) Ifc() interface{} {
	panic("method not available")
}

func (p ParsedDefault) IsValid() bool {
	panic("method not available")
}

type ParsedBool struct {
	Value bool
	Valid bool
}

func (ParsedBool) I32() int32 {
	panic("only bool available")
}

func (ParsedBool) I64() int64 {
	panic("only bool available")
}

func (ParsedBool) F32() float32 {
	panic("only bool available")
}

func (ParsedBool) F64() float64 {
	panic("only bool available")
}

func (ParsedBool) String() string {
	panic("only bool available")
}

func (p ParsedBool) Bool() bool {
	return p.Value
}

func (ParsedBool) Time() time.Time {
	panic("only bool available")
}

func (ParsedBool) Ifc() interface{} {
	panic("only bool available")
}

func (p ParsedBool) IsValid() bool {
	return p.Valid
}

type ParsedInt32 int32

func (p ParsedInt32) I32() int32 {
	return int32(p)
}

func (ParsedInt32) I64() int64 {
	panic("only int32 available")
}

func (ParsedInt32) F32() float32 {
	panic("only int32 available")
}

func (ParsedInt32) F64() float64 {
	panic("only int32 available")
}

func (ParsedInt32) String() string {
	panic("only int32 available")
}

func (ParsedInt32) Bool() bool {
	panic("only int32 available")
}

func (ParsedInt32) Time() time.Time {
	panic("only int32 available")
}

func (ParsedInt32) Ifc() interface{} {
	panic("only int32 available")
}

func (ParsedInt32) IsValid() bool {
	return true
}

type ParsedInt64 int64

func (ParsedInt64) I32() int32 {
	panic("only int64 available")
}

func (p ParsedInt64) I64() int64 {
	return int64(p)
}

func (ParsedInt64) F32() float32 {
	panic("only int64 available")
}

func (ParsedInt64) F64() float64 {
	panic("only int64 available")
}

func (ParsedInt64) String() string {
	panic("only int64 available")
}

func (ParsedInt64) Bool() bool {
	panic("only int64 available")
}

func (ParsedInt64) Time() time.Time {
	panic("only int64 available")
}

func (ParsedInt64) Ifc() interface{} {
	panic("only int64 available")
}

func (ParsedInt64) IsValid() bool {
	return true
}

type ParsedFloat32 float32

func (ParsedFloat32) I32() int32 {
	panic("only float32 available")
}

func (ParsedFloat32) I64() int64 {
	panic("only float32 available")
}

func (p ParsedFloat32) F32() float32 {
	return float32(p)
}

func (ParsedFloat32) F64() float64 {
	panic("only float32 available")
}

func (ParsedFloat32) String() string {
	panic("only float32 available")
}

func (ParsedFloat32) Bool() bool {
	panic("only float32 available")
}

func (ParsedFloat32) Time() time.Time {
	panic("only float32 available")
}

func (ParsedFloat32) Ifc() interface{} {
	panic("only float32 available")
}

func (ParsedFloat32) IsValid() bool {
	return true
}

var _ ParsedValue = ParsedFloat32(0)

type ParsedFloat64 float64

func (ParsedFloat64) I32() int32 {
	panic("only float64 available")
}

func (ParsedFloat64) I64() int64 {
	panic("only float64 available")
}

func (ParsedFloat64) F32() float32 {
	panic("only float64 available")
}

func (p ParsedFloat64) F64() float64 {
	return float64(p)
}

func (ParsedFloat64) String() string {
	panic("only float64 available")
}

func (ParsedFloat64) Bool() bool {
	panic("only float64 available")
}

func (ParsedFloat64) Time() time.Time {
	panic("only float64 available")
}

func (ParsedFloat64) Ifc() interface{} {
	panic("only float64 available")
}

func (ParsedFloat64) IsValid() bool {
	return true
}

type ParsedString string

func (ParsedString) I32() int32 {
	panic("only string available")
}

func (ParsedString) I64() int64 {
	panic("only string available")
}

func (ParsedString) F32() float32 {
	panic("only string available")
}

func (ParsedString) F64() float64 {
	panic("only string available")
}

func (p ParsedString) String() string {
	return string(p)
}

func (ParsedString) Bool() bool {
	panic("only string available")
}

func (ParsedString) Time() time.Time {
	panic("only string available")
}

func (ParsedString) Ifc() interface{} {
	panic("only string available")
}

func (ParsedString) IsValid() bool {
	return true
}

type ParsedTime time.Time

func (ParsedTime) I32() int32 {
	panic("only time.Time available")
}

func (ParsedTime) I64() int64 {
	panic("only time.Time available")
}

func (ParsedTime) F32() float32 {
	panic("only time.Time available")
}

func (ParsedTime) F64() float64 {
	panic("only time.Time available")
}

func (ParsedTime) String() string {
	panic("only time.Time available")
}

func (ParsedTime) Bool() bool {
	panic("only time.Time available")
}

func (p ParsedTime) Time() time.Time {
	return time.Time(p)
}

func (ParsedTime) Ifc() interface{} {
	panic("only time.Time available")
}

func (ParsedTime) IsValid() bool {
	return true
}

type ParsedDecimal struct {
	ParsedDefault
	b *decimal.Big
}

func (p ParsedDecimal) Ifc() interface{} {
	return p.b
}

func (p ParsedDecimal) IsValid() bool {
	return p.b != nil
}
