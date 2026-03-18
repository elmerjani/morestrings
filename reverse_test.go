package morestrings

import (
	"fmt"
	"testing"
)

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		in   string //in
		want string // want
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
	var x int
	x = 12
	
	if x < 10 {
		fmt.Printf("x is less than 10")
	}
}
