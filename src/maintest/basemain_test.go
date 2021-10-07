package maintest

import (
    "testing"
)

var intVar int
var stringVar string

//go:noinline
func emptyInterfaceInt(i interface {}) {
    str := i.(int)
    intVar = str
}

//go:noinline
func withTypeInt(str int) {
    intVar = str
}

//go:noinline
func emptyInterfaceString(i interface {}) {
    str := i.(string)
    stringVar = str
}

//go:noinline
func withTypeString(str string) {
    stringVar = str
}

func BenchmarkWithTypeInt(b *testing.B) {
    str := 123
    for i := 0; i < b.N; i++ {
        withTypeInt(str)
    }
}

func BenchmarkWithEmptyInterfaceInt(b *testing.B) {
    str := 123
    for i := 0; i < b.N; i++ {
        emptyInterfaceInt(str)
    }
}

func BenchmarkWithTypeString(b *testing.B) {
    str := "123"
    for i := 0; i < b.N; i++ {
        withTypeString(str)
    }
}

func BenchmarkWithEmptyInterfaceString(b *testing.B) {
    str := "123"
    for i := 0; i < b.N; i++ {
        emptyInterfaceString(str)
    }
}
