// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"testing"
)

type BenchmarkStruct struct {
	Name    string
	Age     int
	Active  bool
	Score   float32
	Rank    float64
	Tags    []string
	Pointer *int
}

func BenchmarkStringify(b *testing.B) {
	val := 42
	s := &BenchmarkStruct{
		Name:    "benchmark",
		Age:     30,
		Active:  true,
		Score:   1.1,
		Rank:    99.999999,
		Tags:    []string{"go", "github", "api"},
		Pointer: &val,
	}
	b.ResetTimer()
	for range b.N { //nolint:modernize // b.Loop() requires Go 1.24
		Stringify(s)
	}
}

func TestStringify_Floats(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in  any
		out string
	}{
		{float32(1.1), "1.1"},
		{float64(1.1), "1.1"},
		{float32(1.0000001), "1.0000001"},
		{struct{ F float32 }{1.1}, "{F:1.1}"},
	}

	for i, tt := range tests {
		s := Stringify(tt.in)
		if s != tt.out {
			t.Errorf("%v. Stringify(%v) = %q, want %q", i, tt.in, s, tt.out)
		}
	}
}
