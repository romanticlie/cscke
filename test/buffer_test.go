package test

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkBuilder(b *testing.B) {

	var s strings.Builder

	s.Grow(10e7)

	for i := 0;i < 10e7;i++ {
		s.WriteString("1")
	}
}

func BenchmarkBuffer(b *testing.B){

	var s bytes.Buffer

	s.Grow(10e7)

	for i := 0;i < 10e7;i++ {
		s.WriteString("1")
	}
}


