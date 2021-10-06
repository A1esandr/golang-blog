package maintest

import (
	"testing"
)

var instance *StructWithFields

type StructWithFields struct {
	field1 int
	field2 string
	field3 float32
	field4 float64
	field5 int32
	field6 bool
	field7 uint64
	field8 *string
	field9 uint16
}

//go:noinline
func emptyInterface(i interface {}) {
	str := i.(*StructWithFields)
	instance = str
}

//go:noinline
func withType(str *StructWithFields) {
	instance = str
}

func BenchmarkWithType(b *testing.B) {
	str := StructWithFields{field1: 1, field2: "string", field8: new(string)}
	for i := 0; i < b.N; i++ {
		withType(&str)
	}
}

func BenchmarkWithEmptyInterface(b *testing.B) {
	str := StructWithFields{field1: 1, field2: "string", field8: new(string)}
	for i := 0; i < b.N; i++ {
		emptyInterface(&str)
	}
}
